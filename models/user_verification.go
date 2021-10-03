package models

import "time"

type UserVerification struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	Type         string    `db:"type"`        // it will be 'register' and 'forgot_password' or 'phone_verification'
	Medium       string    `db:"medium"`      // it will be 'email' or 'phone'
	Destination  string    `db:"destination"` // it will be phone_number or email user
	RequestCount int       `db:"request_count"`
	SubmitCount  int       `db:"submit_count"`
	Deeplink     string    `db:"deeplink"`
	OTP          string    `db:"string"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (uv *UserVerification) GetId() int {
	return uv.ID
}

func (uv *UserVerification) SetUserId(userId int) {
	uv.UserID = userId
}

func (uv *UserVerification) GetUserId() int {
	return uv.UserID
}

func (uv *UserVerification) SetType(types string) {
	uv.Type = types
}

func (uv *UserVerification) GetType() string {
	return uv.Type
}

func (uv *UserVerification) SetMedium(medium string) {
	uv.Medium = medium
}

func (uv *UserVerification) GetMedium() string {
	return uv.Medium
}

func (uv *UserVerification) SetDestination(destination string) {
	uv.Destination = destination
}

func (uv *UserVerification) GetDestination() string {
	return uv.Destination
}

func (uv *UserVerification) SetDeeplink(deeplink string) {
	uv.Deeplink = deeplink
}

func (uv *UserVerification) GetDeeplink() string {
	return uv.Deeplink
}

func (uv *UserVerification) SetOTP(otp string) {
	uv.OTP = otp
}

func (uv *UserVerification) GetOTP() string {
	return uv.OTP
}

func (uv *UserVerification) SetRequestCount(requestCount int) {
	uv.RequestCount = requestCount
}

func (uv *UserVerification) GetRequestCount() int {
	return uv.RequestCount
}

func (uv *UserVerification) SetSubmitCount(submitCount int) {
	uv.SubmitCount = submitCount
}

func (uv *UserVerification) GetSubmitCount() int {
	return uv.SubmitCount
}

func (uv *UserVerification) SetCreatedAt(createdAt time.Time) {
	uv.CreatedAt = createdAt
}

func (uv *UserVerification) GetCreatedAt() time.Time {
	return uv.CreatedAt
}

func (uv *UserVerification) SetUpdatedAt(updatedAt time.Time) {
	uv.UpdatedAt = updatedAt
}

func (uv *UserVerification) GetUpdatedAt() time.Time {
	return uv.UpdatedAt
}
