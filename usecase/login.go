package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"reflect"
	"sync"

	"spos/auth/constant"
	"spos/auth/models"
	"spos/auth/models/view"

	"github.com/s-pos/go-utils/utils/response"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(ctx context.Context, req models.RequestLogin) response.Output {
	var (
		errChan = make(chan error, 1)
		wg      sync.WaitGroup
	)

	user, err := u.repository.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.Errors(
				ctx,
				http.StatusNotFound,
				string(constant.UserNotFound),
				constant.Message[constant.LoginFailed],
				constant.Reason[constant.UserNotFound],
				err,
			)
		}
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.ErrorQueryFind), constant.Message[constant.LoginFailed], constant.ErrorGlobal, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password)); err != nil {
		return response.Errors(ctx, http.StatusBadRequest, string(constant.UserPasswordNotMatch), constant.Message[constant.LoginFailed], constant.Reason[constant.UserPasswordNotMatch], err)
	}

	user.SetFcmToken(req.FcmToken)
	user.SetDeviceId(req.DeviceID)
	accessToken, expIn, err := u.repository.SetAccessToken(ctx, user)
	if err != nil {
		var code = string(constant.ErrorRedisSet)
		if errors.Is(err, errors.New(string(constant.ErrorMarshal))) {
			code = string(constant.ErrorMarshal)
		}

		return response.Errors(ctx, http.StatusInternalServerError, code, constant.Message[constant.LoginFailed], constant.ErrorGlobal, err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		if !reflect.ValueOf(req.DeviceID).IsZero() && !reflect.ValueOf(req.FcmToken).IsZero() {
			_, err = u.repository.InsertFcmToken(user)
			if err != nil {
				errChan <- err
			}
		}
	}()

	loginView := view.LoginView{
		AccessToken: accessToken,
		TokenType:   constant.BearerToken,
		ExpiresIn:   expIn,
	}
	wg.Wait()
	close(errChan)

	return response.Success(ctx, http.StatusOK, string(constant.LoginSuccess), constant.Message[constant.LoginSuccess], loginView)
}
