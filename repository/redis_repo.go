package repository

import (
	"context"
	"encoding/json"
	"errors"
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
		"id":           user.Id(),
		"name":         user.Name(),
		"email":        user.Email(),
		"phone_number": user.Phone(),
		"fcm_token":    user.FcmToken(),
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		logger.Messagef("error marshal %v", err).To(ctx)
		return key, int64(exp.Seconds()), errors.New(string(constant.ErrorMarshal))
	}

	result := r.redis.Set(ctx, key, string(dataByte), exp)
	return key, int64(exp.Seconds()), result.Err()
}

func (r *repo) GetRedisData(ctx context.Context, key string) (string, error) {
	result := r.redis.Get(ctx, key)

	return result.Val(), result.Err()
}
