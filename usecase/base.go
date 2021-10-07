package usecase

import (
	"context"
	"time"

	"spos/auth/models"
	"spos/auth/repository"
	"spos/auth/usecase/rpc"

	"github.com/s-pos/go-utils/utils/response"
)

type usecase struct {
	authClient rpc.AuthClient
	repository repository.Repository
	location   *time.Location
}

type Usecase interface {
	// Login logic for user doing login
	Login(ctx context.Context, req models.RequestLogin) response.Output

	// Register logic for user doing register
	Register(ctx context.Context, req models.RequestRegister) response.Output
}

func New(rpc rpc.AuthClient, repo repository.Repository, loc *time.Location) Usecase {
	return &usecase{
		authClient: rpc,
		repository: repo,
		location:   loc,
	}
}
