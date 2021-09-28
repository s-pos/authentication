package usecase

import (
	"context"
	"time"

	"spos/auth/models"
	"spos/auth/repository"

	"github.com/s-pos/go-utils/utils/response"
)

type usecase struct {
	repository repository.Repository
	location   *time.Location
}

type Usecase interface {
	// Login logic for user doing login
	Login(ctx context.Context, req models.RequestLogin) response.Output
}

func New(repo repository.Repository, loc *time.Location) Usecase {
	return &usecase{
		repository: repo,
		location:   loc,
	}
}
