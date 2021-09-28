package models

type RequestLogin struct {
	Email    string `json:"email" required:"json"`
	Password string `json:"password" required:"json"`
	FcmToken string `json:"fcmToken"`
	DeviceID string `json:"deviceId"`
}
