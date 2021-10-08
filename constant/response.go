package constant

type Code string

const (
	// ErrorGlobal for global error message
	ErrorGlobal = "Terjadi kesalahan, silahkan coba beberapa saat lagi"
	// RegisterSuccessMessage is message response when user success register
	RegisterSuccessMessage = "Pendaftaran berhasil. silakan cek email %s untuk melanjutkan verifikasi"

	// LoginSuccess code for success login
	LoginSuccess Code = "100010"
	// LoginFailed code for failed login
	LoginFailed Code = "100090"

	// RegisterSuccess code for success register
	RegisterSuccess Code = "101010"
	// RegisterFailed code for failed register
	RegisterFailed Code = "101090"

	// VerificationSuccess code for success verification register with otp
	VerificationSuccess Code = "101110"
	// VerificationFailed code for failed verification register
	VerificationFailed Code = "101190"
	// VerificationUserFailed code for failed on userVerification
	// query find or updated
	VerificationUserFailed Code = "101191"

	// UserNotFound when user try to login but email not found from database
	UserNotFound Code = "108140"
	// UserPasswordNotMatch password not match from database with request payload
	UserPasswordNotMatch Code = "108141"
	// UserNotVerified user not yet verified email
	UserNotVerified Code = "108142"
	// UserEmailAlreadyUsed when user register, but email already on database
	UserEmailAlreadyUsed Code = "108143"
	// UserPhoneAlreadyUsed when user register, but phone already on database
	UserPhoneAlreadyUsed Code = "108144"
	// UserAlreadyRequestOTP interval 2 minutes for another request OTP
	UserAlreadyRequestOTP Code = "108145"
	// UserEmailNotSame email from redis and from request not same
	// during verification register
	UserEmailNotSame Code = "108146"
	// OTPInvalid otp not found or expired in redis
	OTPInvalid Code = "108540"

	/* ========== GLOBAL ERROR WILL BE HERE ========== */

	// ErrorQueryFind is error code for query find/get data
	ErrorQueryFind Code = "109090"
	// ErrorQueryUpdate is error code for query update data
	ErrorQueryUpdate Code = "109091"
	// ErrorQueryInsert is error code for query insert data
	ErrorQueryInsert Code = "109092"

	// ErrorRedisSet is error code for set redis
	ErrorRedisSet Code = "109190"
	// ErrorRedisGet is error code for get data from redis
	ErrorRedisGet Code = "109191"

	// BodyRequired is when request body required
	BodyRequired Code = "109290"

	// ErrorMarshal error when marshal
	ErrorMarshal Code = "109990"
	// ErrorUnmarshal error when unmarshal struct
	ErrorUnmarshal Code = "109991"
)

var (
	Message = map[Code]string{
		LoginSuccess: "login.success",
		LoginFailed:  "login.failed",

		RegisterSuccess: "register.success",
		RegisterFailed:  "register.failed",

		VerificationSuccess: "verification.success",
		VerificationFailed:  "verification.failed",

		UserAlreadyRequestOTP: "request.failed",
	}

	Reason = map[Code]string{
		VerificationSuccess:   "Verifikasi berhasil, silakan login",
		UserNotFound:          "User tidak ditemukan",
		UserPasswordNotMatch:  "Email atau Password tidak sesuai",
		UserNotVerified:       "Anda belum melakukan verifikasi, silakan verifikasi diri Anda terlebih dahulu",
		UserEmailAlreadyUsed:  "Email telah digunakan, silakan gunakan email lain",
		UserPhoneAlreadyUsed:  "Nomor telepon telah digunakan, silakan gunakan nomor yang lain",
		UserAlreadyRequestOTP: "Anda baru saja melakukan permintaan pengiriman OTP, tunggu beberapa saat lagi",
		UserEmailNotSame:      "Permintaan kode verifikasi tidak valid",
		OTPInvalid:            "Kode verifikasi tidak ditemukan atau sudah tidak berlaku",

		BodyRequired: "Permintaan tidak lengkap, silakan cek kembali",
	}
)
