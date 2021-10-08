package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"spos/auth/constant"
	"spos/auth/models"

	"github.com/s-pos/go-utils/logger"
	"github.com/s-pos/go-utils/utils/response"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, req models.RequestRegister) response.Output {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 1)
	)
	user, err := u.repository.GetUserByEmail(req.Email)
	if err == nil {
		// if email already verified
		// send error response email already used
		if user.IsEmailVerified() {
			err = fmt.Errorf("email already used")
			return response.Errors(ctx, http.StatusBadRequest, string(constant.UserEmailAlreadyUsed), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserEmailAlreadyUsed], err)
		}

		// checking if email not yet verified but password is wrong
		// then send error response email already used
		// that means that user want to regist but wrong password (like login)
		if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password)); err != nil {
			err = fmt.Errorf("email already used")
			return response.Errors(ctx, http.StatusBadRequest, string(constant.UserEmailAlreadyUsed), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserEmailAlreadyUsed], err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			resAuthClient, err := u.authClient.SendEmailVerification(ctx, user)
			if err != nil {
				errChan <- err
			}
			logger.Messagef("%v", resAuthClient).To(ctx)
		}()

		wg.Wait()
		select {
		case err = <-errChan:
			return response.Errors(ctx, http.StatusInternalServerError, string(constant.UserAlreadyRequestOTP), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserAlreadyRequestOTP], err)
		default:
			close(errChan)
		}
		return response.Success(ctx, http.StatusOK, string(constant.RegisterSuccess), constant.Message[constant.RegisterSuccess], fmt.Sprintf(constant.RegisterSuccessMessage, user.GetEmail()))
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.ErrorQueryFind), constant.Message[constant.RegisterFailed], constant.ErrorGlobal, err)
	}

	// set no telp and will return 628xxxx
	user.SetPhone(req.PhoneNumber)
	user, err = u.repository.GetUserByPhone(user.GetPhone())
	if err == nil {
		err = fmt.Errorf("phone number already used")
		return response.Errors(ctx, http.StatusBadRequest, string(constant.UserPhoneAlreadyUsed), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserPhoneAlreadyUsed], err)
	}

	passwordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.ErrorUnmarshal), constant.Message[constant.RegisterFailed], constant.ErrorGlobal, err)
	}
	user.SetPhone(req.PhoneNumber)
	user.SetPassword(string(passwordByte))
	user.SetName(req.Name)
	user.SetEmail(req.Email)

	user, err = u.repository.InsertNewUser(user)
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, string(constant.ErrorQueryInsert), constant.Message[constant.RegisterFailed], constant.ErrorGlobal, err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		resAuthClient, err := u.authClient.SendEmailVerification(ctx, user)
		if err != nil {
			errChan <- err
		}
		logger.Messagef("%v", resAuthClient).To(ctx)
	}()

	wg.Wait()
	select {
	case err = <-errChan:
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.UserAlreadyRequestOTP), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserAlreadyRequestOTP], err)
	default:
		close(errChan)
	}

	return response.Success(ctx, http.StatusCreated, string(constant.RegisterSuccess), constant.Message[constant.RegisterSuccess], fmt.Sprintf(constant.RegisterSuccessMessage, user.GetEmail()))
}
