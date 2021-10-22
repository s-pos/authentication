package rpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"

	"spos/auth/models"
	"spos/auth/repository"

	"github.com/s-pos/protobuf/go/auth"
	"google.golang.org/grpc"
)

type authClient struct {
	client     auth.UserAuthServiceClient
	httpClient *http.Client
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
		emailGrpcPort = os.Getenv("QUEUE_GRPC_HOST")
		opts          []grpc.DialOption
	)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(emailGrpcPort, opts...)
	if err != nil {
		err = fmt.Errorf("error connect to grpc client %s and got error %v", emailGrpcPort, err)
		panic(err)
	}

	client := auth.NewUserAuthServiceClient(conn)

	return &authClient{
		client:     client,
		repository: repo,
		timezone:   timezone,
		httpClient: httpClient,
	}
}
