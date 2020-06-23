package e

var (
	// Common errors
	Ok                  = &E{Code: 0, Message: "OK"}
	InternalServerError = &E{Code: 10001, Message: "Internal server error"}
	ErrBind             = &E{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &E{Code: 10003, Message: "Param error"}

	// client errors
	ParamPhoneEmpty   = &E{Code: 20001, Message: "param phone is empty"}
	GetPhoneCodeOffen = &E{Code: 20002, Message: "get phone code offen"}
	ParamCodeEmpty    = &E{Code: 20003, Message: "param code is empty"}
)
