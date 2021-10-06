package repository

import (
	"log"
	"time"

	"spos/auth/models"
)

func (r *repo) GetUserByEmail(email string) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	query := `select
       		id, "name", email, phone_number, password
			from users
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
       		email_verification_at
			from users
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
		user.GetPassword(), now, now,
	).StructScan(user)

	if err != nil {
		tx.Rollback()
		return user, err
	}
	log.Println(user)
	err = tx.Commit()
	return user, err
}
