package repository

import "spos/auth/models"

func (r *repo) GetUserVerification(userId int, medium, dest string) (*models.UserVerification, error) {
	var (
		user models.UserVerification
		err  error
	)
	query := `select
			id, user_id, "type", request_count,
			submit_count, updated_at, deeplink, otp
			from user_verifications
			where user_id=$1 and medium=$2 and destination=$3`

	err = r.db.Get(&user, query, userId, medium, dest)
	return &user, err
}

func (r *repo) NewUserVerification(userVerification *models.UserVerification) (*models.UserVerification, error) {
	var (
		uv  models.UserVerification
		err error
		tx  = r.db.MustBegin()
	)
	query := `insert into user_verifications
			(user_id, type, medium, destination, request_count, deeplink, otp, submit_count, created_at, updated_at)
			values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id`
	err = tx.QueryRowx(
		query, userVerification.GetUserId(),
		userVerification.GetType(), userVerification.GetMedium(),
		userVerification.GetDestination(), userVerification.GetRequestCount(),
		userVerification.GetDeeplink(), userVerification.GetOTP(),
		userVerification.GetSubmitCount(), userVerification.GetCreatedAt(),
		userVerification.GetUpdatedAt(),
	).StructScan(&uv)
	if err != nil {
		tx.Rollback()
		return userVerification, err
	}

	err = tx.Commit()
	return &uv, err
}

func (r *repo) UpdateUserVerification(userVerification *models.UserVerification) (*models.UserVerification, error) {
	var (
		uv  models.UserVerification
		err error
		tx  = r.db.MustBegin()
	)
	query := `update user_verifications set
			request_count=$1, submit_count=$2, updated_at=$3,
            deeplink=$4, otp=$5
			where user_id=$6 and destination=$7 and medium=$8`
	err = tx.QueryRowx(
		query, userVerification.GetRequestCount(),
		userVerification.GetSubmitCount(), userVerification.GetUpdatedAt(),
		userVerification.GetDeeplink(), userVerification.GetOTP(),
		userVerification.GetId(), userVerification.GetDestination(),
		userVerification.GetMedium(),
	).StructScan(&uv)
	if err != nil {
		tx.Rollback()
		return userVerification, err
	}

	err = tx.Commit()
	return &uv, err
}
