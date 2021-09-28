package constant

type Code string

const (
	// ErrorGlobal for global error message
	ErrorGlobal = "Terjadi kesalahan, silahkan coba beberapa saat lagi"

	// LoginSuccess code for success login
	LoginSuccess Code = "100010"
	// LoginFailed code for failed login
	LoginFailed Code = "100090"

	// UserNotFound when user try to login but email not found from database
	UserNotFound Code = "108140"
	// UserPasswordNotMatch password not match from database with request payload
	UserPasswordNotMatch Code = "108141"
	// UserNotVerified user not yet verified email
	UserNotVerified Code = "108142"

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
	}

	Reason = map[Code]string{
		UserNotFound:         "User tidak ditemukan",
		UserPasswordNotMatch: "Email atau Password tidak sesuai",
		UserNotVerified:      "Anda belum melakukan verifikasi, silahkan verifikasi diri Anda terlebih dahulu",

		BodyRequired: "Permintaan tidak lengkap, silahkan cek kembali",
	}
)
