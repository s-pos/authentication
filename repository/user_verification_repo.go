package repository

import (
	"spos/auth/models"
)

func (r *repo) GetUserVerificationByDestination(medium, dest string) (*models.UserVerification, error) {
	var (
		uv  models.UserVerification
		err error
	)
	query := `select
			id, user_id, "type", request_count,
			submit_count, updated_at, deeplink, otp,
       		medium, destination, created_at
			from user_verifications
			where medium=$1 and destination=$2`

	err = r.db.Get(&uv, query, medium, dest)
	return &uv, err
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
		userVerification.GetSubmitCount(), userVerification.GetCreatedAt().UTC(),
		userVerification.GetUpdatedAt().UTC(),
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
						where id=$6 returning id`
	err = tx.QueryRowx(
		query, userVerification.GetRequestCount(),
		userVerification.GetSubmitCount(), userVerification.GetUpdatedAt().UTC(),
		userVerification.GetDeeplink(), userVerification.GetOTP(),
		userVerification.GetId(),
	).StructScan(&uv)
	if err != nil {
		tx.Rollback()
		return userVerification, err
	}

	err = tx.Commit()
	return &uv, err
}
