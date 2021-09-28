package models

type RequestLogin struct {
	Email    string `json:"email" required:"json"`
	Password string `json:"password" required:"json"`
	FcmToken string `json:"fcm_token"`
	DeviceID string `json:"device_id"`
}
