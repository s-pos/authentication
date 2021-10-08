package rpc

import (
	"context"
	"database/sql"
	"errors"
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

	err = ac.saveUserVerification(user, constant.MediumEmail, constant.TypeRegister, otp)
	if err != nil {
		logger.Messagef("error save verification %v", err).To(ctx)
		return nil, err
	}

	// req := &auth.Verification{
	// 	Name:  user.GetName(),
	// 	Email: user.GetEmail(),
	// 	Otp:   otp,
	// }

	// res, err := ac.client.SendEmailVerification(ctx, req)

	return nil, err
}

func (ac *authClient) saveUserVerification(user *models.User, medium, types, otp string) error {
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
		return err
	}

	uv, err := ac.repository.GetUserVerificationByDestination(medium, dest)
	if err == nil {
		if uv.IsReadyToSend() {
			return nil
		}

		err = fmt.Errorf("you just send request verification. please wait 2mins")
		return err
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	uv.SetUserId(user.GetId())
	uv.SetOTP(otp)
	uv.SetDestination(dest)
	uv.SetRequestCount(1)
	uv.SetType(types)
	uv.SetMedium(medium)
	uv.SetCreatedAt(now)

	_, err = ac.repository.NewUserVerification(uv)
	if err != nil {
		return err
	}

	return nil
}
