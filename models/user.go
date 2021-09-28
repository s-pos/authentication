package models

import "time"

type User struct {
	id                  int        `db:"id"`
	name                string     `db:"name"`
	email               string     `db:"email"`
	password            string     `db:"password"`
	phone               string     `db:"phone"`
	emailVerificationAt *time.Time `db:"email_verification_at"`
	phoneVerificationAt *time.Time `db:"phone_verification_at"`
	createdAt           time.Time  `db:"created_at"`
	updatedAt           time.Time  `db:"updated_at"`

	fcmToken string `db:"-"`
}

func (u *User) FcmToken() string {
	return u.fcmToken
}

func (u *User) SetFcmToken(fcmToken string) {
	u.fcmToken = fcmToken
}

func NewUser() *User {
	return &User{}
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Phone() string {
	return u.phone
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u *User) EmailVerificationAt() *time.Time {
	return u.emailVerificationAt
}

func (u *User) SetEmailVerificationAt(emailVerificationAt *time.Time) {
	u.emailVerificationAt = emailVerificationAt
}

func (u *User) PhoneVerificationAt() *time.Time {
	return u.phoneVerificationAt
}

func (u *User) SetPhoneVerificationAt(phoneVerificationAt *time.Time) {
	u.phoneVerificationAt = phoneVerificationAt
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}
