package repository

import (
	"time"

	"spos/auth/models"
)

func (r *repo) InsertFcmToken(user *models.User) (*models.FcmToken, error) {
	var (
		fcmToken models.FcmToken
		now      = time.Now().In(r.location)
		tx       = r.db.MustBegin()
		err      error
	)
	query := `insert into fcm_tokens
			(user_id, token, device_id, created_at)
			values($1, $2, $3, $4)
			returning id, user_id, token, device_id`
	err = tx.QueryRowx(
		query,
		user.GetId(),
		user.GetFcmToken(),
		user.GetDeviceId(),
		now,
	).StructScan(&fcmToken)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return &fcmToken, err
}
