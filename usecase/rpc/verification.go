package rpc

import (
	"context"
	"spos/auth/models"

	"github.com/s-pos/protobuf/go/auth"
)

func (ac *authClient) SendEmailVerification(ctx context.Context, user *models.User) (*auth.VerificationReply, error) {
	otp, err := ac.repository.SetToken(ctx, user.GetEmail(), false)
	if err != nil {
		return nil, err
	}

	req := &auth.Verification{
		Name:  user.GetName(),
		Email: user.GetEmail(),
		Otp:   otp,
	}

	res, err := ac.client.SendEmailVerification(ctx, req)

	return res, err
}
