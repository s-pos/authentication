package models

import "time"

type FcmToken struct {
	id        int       `db:"id"`
	userId    int       `db:"user_id"`
	token     string    `db:"token"`
	deviceId  *string   `db:"deviceId"`
	createdAt time.Time `db:"created_at"`
}

func (f *FcmToken) Id() int {
	return f.id
}

func (f *FcmToken) SetId(id int) {
	f.id = id
}

func (f *FcmToken) UserId() int {
	return f.userId
}

func (f *FcmToken) SetUserId(userId int) {
	f.userId = userId
}

func (f *FcmToken) Token() string {
	return f.token
}

func (f *FcmToken) SetToken(token string) {
	f.token = token
}

func (f *FcmToken) DeviceId() string {
	if f.deviceId != nil {
		return *f.deviceId
	}
	return ""
}

func (f *FcmToken) SetDeviceId(deviceId string) {
	f.deviceId = &deviceId
}

func (f *FcmToken) CreatedAt() time.Time {
	return f.createdAt
}

func (f *FcmToken) SetCreatedAt(createdAt time.Time) {
	f.createdAt = createdAt
}

func NewFcmToken() *FcmToken {
	return &FcmToken{}
}
