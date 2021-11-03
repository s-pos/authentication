package repository

import (
	"time"

	"spos/auth/models"
)

func (r *repo) GetUserByEmail(email string) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	query := `select
      id, "name", email, phone_number, password, 
			email_verification_at, phone_verification_at,
			user_verifications
			from users_view
			where email=$1`

	err = r.db.Get(&user, query, email)
	return &user, err
}

func (r *repo) GetUserByPhone(phone string) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	query := `select
      id, "name", email, phone_number, password,
      email_verification_at, phone_verification_at,
			user_verifications
			from users_view
			where phone_number=$1`

	err = r.db.Get(&user, query, phone)
	return &user, err
}

func (r *repo) InsertNewUser(user *models.User) (*models.User, error) {
	var (
		tx  = r.db.MustBegin()
		now = time.Now().In(r.location)
		err error
	)
	query := `insert into users
			(name, email, phone_number, password, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6)
			returning id`

	err = tx.QueryRowx(
		query, user.GetName(),
		user.GetEmail(), user.GetPhone(),
		user.GetPassword(), now.UTC(), now.UTC(),
	).StructScan(user)

	if err != nil {
		tx.Rollback()
		return user, err
	}
	err = tx.Commit()
	return user, err
}

func (r *repo) UpdateUser(user *models.User) (*models.User, error) {
	var (
		tx  = r.db.MustBegin()
		err error
	)
	query := `update users set 
					password=$1, email_verification_at=$2,
					phone_verification_at=$3, updated_at=$4
					where id=$5 returning id`

	err = tx.QueryRowx(
		query,
		user.GetPassword(),
		user.GetEmailVerificationAt().UTC(),
		user.GetPhoneVerificationAt().UTC(),
		user.GetUpdatedAt().UTC(),
		user.GetId(),
	).StructScan(user)

	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	return user, err
}
