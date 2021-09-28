package repository

import (
	"time"

	"spos/auth/models"
)

func (r *repo) GetUserByEmail(email string) (*models.User, error) {
	var (
		user = new(models.User)
		err  error
	)

	query := `select
       		id, "name", email, phone, password,
       		email_verification_at, phone_verification_at
			from users
			where email=$1`

	err = r.db.Get(&user, query, email)
	return user, err
}

func (r *repo) GetUserByPhone(phone string) (*models.User, error) {
	var (
		user = new(models.User)
		err  error
	)

	query := `select
       		id, "name", email, phone, password,
       		email_verification_at, phone_verification_at
			from users
			where phone=$1`

	err = r.db.Get(&user, query, phone)
	return user, err
}

func (r *repo) InsertNewUser(user *models.User) (*models.User, error) {
	var (
		tx  = r.db.MustBegin()
		now = time.Now().In(r.location)
		err error
	)
	query := `insert into users
			(name, email, phone, password, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6)
			returning id`

	err = tx.QueryRowx(
		query, user.Name(),
		user.Email(), user.Phone(),
		user.Password(), now, now,
	).StructScan(&user)

	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	return user, err
}
