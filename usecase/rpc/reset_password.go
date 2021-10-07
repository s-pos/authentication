package rpc

import (
	"context"
	"spos/auth/models"

	"github.com/s-pos/protobuf/go/auth"
)

func (ac *authClient) SendEmailResetPassword(ctx context.Context, user *models.User) (*auth.ResetPasswordReply, error) {
	token, err := ac.repository.SetToken(ctx, user.GetEmail(), true)
	if err != nil {
		return nil, err
	}

	req := &auth.ResetPassword{
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Token: token,
	}

	res, err := ac.client.SendEmailResetPassword(ctx, req)

	return res, err
}
