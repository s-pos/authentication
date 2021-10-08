package repository

import (
	"context"
	"time"

	"spos/auth/models"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db       *sqlx.DB
	redis    *redis.Client
	location *time.Location
}

type Repository interface {
	// GetUserByEmail query for getting user data by email
	GetUserByEmail(email string) (*models.User, error)

	// GetUserByPhone query for getting user data by phone number
	GetUserByPhone(phone string) (*models.User, error)

	// InsertNewUser query for create a.k.a register new user
	InsertNewUser(user *models.User) (*models.User, error)

	// InsertFcmToken will add new token after user success login
	InsertFcmToken(user *models.User) (*models.FcmToken, error)

	// SetAccessToken for login section.
	// this function will be set access token and will store to redis
	// with key uuid string and value user data + fcm
	SetAccessToken(ctx context.Context, user *models.User) (string, int64, error)

	// SetToken using for otp
	SetToken(ctx context.Context, data interface{}, isResetPassword bool) (string, error)

	// GetRedisData for global get data from redis
	GetRedisData(ctx context.Context, key string) (string, error)

	// GetUserVerificationByDestination will return data user verification
	GetUserVerificationByDestination(medium, dest string) (*models.UserVerification, error)

	// NewUserVerification query insert for new user (register)
	NewUserVerification(userVerification *models.UserVerification) (*models.UserVerification, error)

	// UpdateUserVerification query untuk update data deeplink, otp, request_count, e.t.c
	UpdateUserVerification(userVerification *models.UserVerification) (*models.UserVerification, error)
}

func New(db *sqlx.DB, redis *redis.Client, location *time.Location) Repository {
	return &repo{
		db:       db,
		redis:    redis,
		location: location,
	}
}
