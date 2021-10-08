package models

import (
	"os"
	"time"
)

type UserVerification struct {
	ID           int        `db:"id"`
	UserID       int        `db:"user_id"`
	Type         string     `db:"type"`        // it will be 'register' and 'forgot_password' or 'phone_verification'
	Medium       string     `db:"medium"`      // it will be 'email' or 'phone'
	Destination  string     `db:"destination"` // it will be phone_number or email user
	RequestCount int        `db:"request_count"`
	SubmitCount  *int       `db:"submit_count"`
	Deeplink     *string    `db:"deeplink"`
	OTP          *string    `db:"otp"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	SubmitedAt   *time.Time `db:"submited_at"`
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
	uv.Deeplink = &deeplink
}

func (uv *UserVerification) GetDeeplink() string {
	if uv.Deeplink == nil {
		return ""
	}

	return *uv.Deeplink
}

func (uv *UserVerification) SetOTP(otp string) {
	uv.OTP = &otp
}

func (uv *UserVerification) GetOTP() string {
	if uv.OTP == nil {
		return ""
	}

	return *uv.OTP
}

func (uv *UserVerification) SetRequestCount(requestCount int) {
	uv.RequestCount = requestCount
}

func (uv *UserVerification) GetRequestCount() int {
	return uv.RequestCount
}

func (uv *UserVerification) SetSubmitCount(submitCount int) {
	uv.SubmitCount = &submitCount
}

func (uv *UserVerification) GetSubmitCount() int {
	if uv.SubmitCount == nil {
		return 0
	}
	return *uv.SubmitCount
}

func (uv *UserVerification) SetCreatedAt(createdAt time.Time) {
	uv.CreatedAt = createdAt
}

func (uv *UserVerification) GetCreatedAt() time.Time {
	return convertTimezone(uv.CreatedAt)
}

func (uv *UserVerification) SetUpdatedAt(updatedAt time.Time) {
	uv.UpdatedAt = &updatedAt
}

func (uv *UserVerification) GetUpdatedAt() *time.Time {
	if uv.UpdatedAt == nil {
		return nil
	}

	updatedAt := convertTimezone(*uv.UpdatedAt)
	return &updatedAt
}

func (uv *UserVerification) SetSubmitedAt(submitedAt time.Time) {
	uv.SubmitedAt = &submitedAt
}

func (uv *UserVerification) GetSubmitedAt() *time.Time {
	if uv.SubmitedAt == nil {
		return &time.Time{}
	}

	submitedAt := convertTimezone(*uv.SubmitedAt)
	return &submitedAt
}

func (uv *UserVerification) IsReadyToSend() bool {
	var (
		now           = time.Now().In(timezone)
		otpTime       time.Time
		otpExpired, _ = time.ParseDuration(os.Getenv("OTP_EXPIRED"))
	)

	otpTime = uv.GetCreatedAt()
	if uv.RequestCount > 1 {
		if uv.GetUpdatedAt() != nil {
			otpTime = *uv.GetUpdatedAt()
		} else {
			otpTime = uv.GetCreatedAt()
		}
	}
	otpTime = otpTime.Add(otpExpired)

	return now.After(otpTime)
}
