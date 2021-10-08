package rpc

import (
	"context"
	"fmt"
	"os"
	"time"

	"spos/auth/models"
	"spos/auth/repository"

	"github.com/s-pos/protobuf/go/auth"
	"google.golang.org/grpc"
)

type authClient struct {
	client     auth.UserAuthServiceClient
	timezone   *time.Location
	repository repository.Repository
}

type AuthClient interface {
	// SendEmailVerification will send request to service email with connection grpc
	//
	// will return a message or error if email failed to send
	SendEmailVerification(ctx context.Context, user *models.User) (*auth.VerificationReply, error)

	// SendEmailResetPassword will send request to service email with connection grpc
	//
	// will return a message or error if email failed to send
	SendEmailResetPassword(ctx context.Context, user *models.User) (*auth.ResetPasswordReply, error)
}

func NewAuthClient(repo repository.Repository, timezone *time.Location) AuthClient {
	var (
		emailGrpcPort = fmt.Sprintf(":%s", os.Getenv("SERVICE_EMAIL_GRPC_PORT"))
		opts          []grpc.DialOption
	)

	conn, err := grpc.Dial(emailGrpcPort, opts...)
	if err != nil {
		return &authClient{
			repository: repo,
			timezone:   timezone,
		}
	}

	client := auth.NewUserAuthServiceClient(conn)

	return &authClient{
		client:     client,
		repository: repo,
		timezone:   timezone,
	}
}
