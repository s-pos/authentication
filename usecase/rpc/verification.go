package rpc

import (
	"context"
	"fmt"
	"time"

	"spos/auth/constant"
	"spos/auth/models"

	"github.com/s-pos/go-utils/logger"
	"github.com/s-pos/protobuf/go/auth"
)

func (ac *authClient) SendEmailVerification(ctx context.Context, user *models.User) (*auth.VerificationReply, error) {
	otp, err := ac.repository.SetToken(ctx, user.GetEmail(), false)
	if err != nil {
		return nil, err
	}

	shortLink, err := ac.saveUserVerification(ctx, user, constant.MediumEmail, constant.TypeRegister, otp)
	if err != nil {
		logger.Messagef("error save verification %v", err).To(ctx)
		return nil, err
	}

	req := &auth.Verification{
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Otp:   otp,
		Link:  shortLink,
	}

	res, err := ac.client.SendEmailVerification(ctx, req)

	return res, err
}

func (ac *authClient) saveUserVerification(ctx context.Context, user *models.User, medium, types, otp string) (string, error) {
	// check data first, if no one, then insert
	var (
		now  = time.Now().In(ac.timezone)
		dest string
		err  error
	)

	switch medium {
	case constant.MediumEmail:
		dest = user.GetEmail()
	case constant.MediumPhone:
		dest = user.GetPhone()
	default:
		err = fmt.Errorf("medium not found")
		return "", err
	}

	uv := user.GetUserVerificationByMediumAndDestination(medium, dest)
	if *uv != (models.UserVerification{}) {
		if uv.IsReadyToSend() {
			// create short link a.k.a dynamic_link
			shortLink, err := ac.dynamicLink(ctx, user, otp)
			if err != nil {
				return "", err
			}

			uv.SetUpdatedAt(now)
			uv.SetOTP(otp)
			uv.SetDeeplink(shortLink)
			uv.SetRequestCount(uv.GetRequestCount() + 1)

			_, err = ac.repository.UpdateUserVerification(uv)
			if err != nil {
				return "", err
			}

			return shortLink, nil
		}

		err = fmt.Errorf(string(constant.UserAlreadyRequestOTP))
		return "", err
	}

	// create short link a.k.a dynamic_link
	shortLink, err := ac.dynamicLink(ctx, user, otp)
	if err != nil {
		return "", err
	}

	uv.SetUserId(user.GetId())
	uv.SetOTP(otp)
	uv.SetDeeplink(shortLink)
	uv.SetDestination(dest)
	uv.SetRequestCount(1)
	uv.SetType(types)
	uv.SetMedium(medium)
	uv.SetCreatedAt(now)

	_, err = ac.repository.NewUserVerification(uv)
	if err != nil {
		return "", err
	}

	return shortLink, nil
}

func (ac *authClient) dynamicLink(ctx context.Context, user *models.User, otp string) (string, error) {
	data := &DynamicLinkData{
		Email: user.GetEmail(),
		Token: otp,
	}

	shortLink, err := ac.CreateDynamicLink(ctx, data, Verification)
	return shortLink, err
}
