package e

var (
	// Common errors
	Ok                  = &E{Code: 0, Message: "OK"}
	InternalServerError = &E{Code: 1001, Message: "Internal server error"}
	ErrBind             = &E{Code: 1002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &E{Code: 1003, Message: "Param error"}

	// client errors
	ParamPhoneEmpty   = &E{Code: 4001, Message: "param phone is empty"}
	GetPhoneCodeOffen = &E{Code: 4002, Message: "get phone code offen"}
	ParamCodeEmpty    = &E{Code: 4003, Message: "param code is empty"}

	CaptchaGenError  = &E{Code: 4004, Message: "gen captcha error"}
	RedisKeyNotExist = &E{Code: 4005, Message: "redis key not exist"}
	VcodeGetfrequent = &E{Code: 4006, Message: "vcode get frequent"}

	TokenMalformed   = &E{Code: 4007, Message: "token malformed"}
	TokenExpired     = &E{Code: 4008, Message: "token expired"}
	TokenNotValidYet = &E{Code: 4009, Message: "token not valid yet"}
	TokenInvalid     = &E{Code: 40010, Message: "token invalid"}
	TokenNull        = &E{Code: 40011, Message: "need auth token"}

	AdminLoginError    = &E{Code: 40013, Message: "name or password invalid"}
	CaptchaInvalid     = &E{Code: 40014, Message: "captcha invalid"}
	UserLoginError     = &E{Code: 40013, Message: "user login error"}
	GetIndexAdminError = &E{Code: 40015, Message: "get admin items error"}
)
