package models

import (
	"fmt"
	"time"
)

type User struct {
	ID                  int        `db:"id"`
	Name                string     `db:"name"`
	Email               string     `db:"email"`
	Password            string     `db:"password"`
	Phone               string     `db:"phone_number"`
	EmailVerificationAt *time.Time `db:"email_verification_at"`
	PhoneVerificationAt *time.Time `db:"phone_verification_at"`
	CreatedAt           time.Time  `db:"created_at"`
	UpdatedAt           time.Time  `db:"updated_at"`

	FcmToken *string `db:"-"`
	DeviceId *string `db:"-"`
}

func (u *User) GetDeviceId() *string {
	return u.DeviceId
}

func (u *User) SetDeviceId(deviceId string) {
	u.DeviceId = &deviceId
}

func (u *User) GetFcmToken() string {
	return *u.FcmToken
}

func (u *User) SetFcmToken(fcmToken string) {
	u.FcmToken = &fcmToken
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetId() int {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) GetPhone() string {
	return u.Phone
}

func (u *User) SetPhone(phone string) {
	switch {
	case phone[:1] == "0":
		u.Phone = fmt.Sprintf("62%s", phone[1:])
	case phone[:1] == "8":
		u.Phone = fmt.Sprintf("62%s", phone[1:])
	case phone[:2] == "62":
		u.Phone = phone
	}
}

func (u *User) GetEmailVerificationAt() *time.Time {
	return u.EmailVerificationAt
}

func (u *User) SetEmailVerificationAt(emailVerificationAt *time.Time) {
	u.EmailVerificationAt = emailVerificationAt
}

func (u *User) IsEmailVerified() bool {
	return u.GetEmailVerificationAt() != nil
}

func (u *User) GetPhoneVerificationAt() *time.Time {
	return u.PhoneVerificationAt
}

func (u *User) SetPhoneVerificationAt(phoneVerificationAt *time.Time) {
	u.PhoneVerificationAt = phoneVerificationAt
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}
