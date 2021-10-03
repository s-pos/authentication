package usecase

import (
	"context"
	"fmt"
	"net/http"

	"spos/auth/constant"
	"spos/auth/models"

	"github.com/s-pos/go-utils/utils/response"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, req models.RequestRegister) response.Output {
	user, err := u.repository.GetUserByEmail(req.Email)
	if err == nil {
		if user.IsEmailVerified() {
			err = fmt.Errorf("email already used")
			return response.Errors(ctx, http.StatusBadRequest, string(constant.UserEmailAlreadyUsed), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserEmailAlreadyUsed], err)
		}
		err = fmt.Errorf("dibuat error dulu aja")
		return response.Errors(ctx, http.StatusBadRequest, string(constant.UserEmailAlreadyUsed), constant.Message[constant.RegisterFailed], constant.Reason[constant.UserEmailAlreadyUsed], err)
	}

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
	user.SetPassword(string(passwordByte))
	user.SetName(req.Name)
	user.SetEmail(req.Email)

	user, err = u.repository.InsertNewUser(user)
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, string(constant.ErrorQueryInsert), constant.Message[constant.RegisterFailed], constant.ErrorGlobal, err)
	}

	return response.Success(ctx, http.StatusCreated, string(constant.RegisterSuccess), constant.Message[constant.RegisterSuccess], fmt.Sprintf(constant.RegisterSuccessMessage, user.GetEmail()))
}
