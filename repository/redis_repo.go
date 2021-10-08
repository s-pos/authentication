package repository

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"spos/auth/constant"
	"spos/auth/models"

	"github.com/google/uuid"
	"github.com/s-pos/go-utils/logger"
)

func (r *repo) SetAccessToken(ctx context.Context, user *models.User) (string, int64, error) {
	var (
		key    = uuid.New().String()
		exp, _ = time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRED"))
	)
	key = strings.ReplaceAll(key, "-", "")
	data := map[string]interface{}{
		"id":           user.GetId(),
		"name":         user.GetName(),
		"email":        user.GetEmail(),
		"phone_number": user.GetPhone(),
		"fcm_token":    user.GetFcmToken(),
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		logger.Messagef("error marshal %v", err).To(ctx)
		return key, int64(exp.Seconds()), errors.New(string(constant.ErrorMarshal))
	}

	result := r.redis.Set(ctx, key, string(dataByte), exp)
	return key, int64(exp.Seconds()), result.Err()
}

func (r *repo) SetToken(ctx context.Context, data interface{}, isResetPassword bool) (string, error) {
	var (
		key           = uuid.New().String()
		otpExpired, _ = time.ParseDuration(os.Getenv("OTP_EXPIRED"))
		err           error
	)
	key = strings.ReplaceAll(key, "-", "")

	if !isResetPassword {
		max := big.NewInt(999999)
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		key = fmt.Sprintf("%06d", n.Int64())
	}

	err = r.redis.Set(ctx, key, data, otpExpired).Err()

	return key, err
}

func (r *repo) GetRedisData(ctx context.Context, key string) (string, error) {
	result := r.redis.Get(ctx, key)

	return result.Val(), result.Err()
}

func (r *repo) DeleteRedisData(ctx context.Context, key string) error {
	err := r.redis.Del(ctx, key).Err()
	return err
}
