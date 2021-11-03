package usecase

import (
	"context"
	"fmt"
	"net/http"
	"spos/auth/constant"
	"spos/auth/models"
	"sync"
	"time"

	"github.com/s-pos/go-utils/utils/response"
)

func (u *usecase) VerificationRegister(ctx context.Context, req models.RequestVerificationOTP) response.Output {
	var (
		now     = time.Now().In(u.location)
		wg      sync.WaitGroup
		errChan = make(chan error, 1)
	)

	user, err := u.repository.GetUserByEmail(req.Email)
	if err != nil {
		return response.Errors(ctx, http.StatusNotFound, string(constant.UserNotFound), constant.Message[constant.VerificationFailed], constant.Reason[constant.UserNotFound], err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		uv := user.GetUserVerificationByMediumAndDestination(constant.MediumEmail, user.GetEmail())
		if uv == nil {
			err = fmt.Errorf("error userverification null")
			errChan <- err
			return
		}

		uv.SetSubmitCount(uv.GetSubmitCount() + 1)
		uv.SetSubmitedAt(now)

		_, err = u.repository.UpdateUserVerification(uv)
		if err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()
	select {
	case err = <-errChan:
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.VerificationUserFailed), constant.Message[constant.VerificationFailed], constant.ErrorGlobal, err)
	default:
		close(errChan)
	}

	email, err := u.repository.GetRedisData(ctx, req.OTP)
	if err != nil {
		return response.Errors(ctx, http.StatusNotFound, string(constant.ErrorRedisGet), constant.Message[constant.VerificationFailed], constant.Reason[constant.OTPInvalid], err)
	}

	if user.GetEmail() != email {
		err = fmt.Errorf("email not same from redis and from request")
		return response.Errors(ctx, http.StatusBadRequest, string(constant.VerificationFailed), constant.Message[constant.VerificationFailed], constant.Reason[constant.UserEmailNotSame], err)
	}

	user.SetEmailVerificationAt(now)
	user.SetUpdatedAt(now)

	_, err = u.repository.UpdateUser(user)
	if err != nil {
		return response.Errors(ctx, http.StatusInternalServerError, string(constant.ErrorQueryUpdate), constant.Message[constant.VerificationFailed], constant.ErrorGlobal, err)
	}

	go u.repository.DeleteRedisData(ctx, req.OTP)

	return response.Success(ctx, http.StatusOK, string(constant.VerificationSuccess), constant.Message[constant.VerificationSuccess], constant.Reason[constant.VerificationSuccess])
}
