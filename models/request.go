package models

type RequestLogin struct {
	Email    string `json:"email" required:"json"`
	Password string `json:"password" required:"json"`
	FcmToken string `json:"fcm_token"`
	DeviceID string `json:"device_id"`
}

type RequestRegister struct {
	Email       string `json:"email" required:"json"`
	Password    string `json:"password" required:"json"`
	PhoneNumber string `json:"phone_number" required:"json"`
	Name        string `json:"name" required:"json"`
}

type RequestVerificationOTP struct {
	Email string `json:"email" required:"json"`
	OTP   string `json:"otp" required:"json"`
}
