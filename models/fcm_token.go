package models

import "time"

type FcmToken struct {
	ID        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Token     string    `db:"token"`
	DeviceId  *string   `db:"device_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (f *FcmToken) Id() int {
	return f.ID
}

func (f *FcmToken) SetId(id int) {
	f.ID = id
}

func (f *FcmToken) GetUserId() int {
	return f.UserId
}

func (f *FcmToken) SetUserId(userId int) {
	f.UserId = userId
}

func (f *FcmToken) GetToken() string {
	return f.Token
}

func (f *FcmToken) SetToken(token string) {
	f.Token = token
}

func (f *FcmToken) GetDeviceId() string {
	if f.DeviceId != nil {
		return *f.DeviceId
	}
	return ""
}

func (f *FcmToken) SetDeviceId(deviceId string) {
	f.DeviceId = &deviceId
}

func (f *FcmToken) GetCreatedAt() time.Time {
	return f.CreatedAt
}

func (f *FcmToken) SetCreatedAt(createdAt time.Time) {
	f.CreatedAt = createdAt
}

func NewFcmToken() *FcmToken {
	return &FcmToken{}
}
