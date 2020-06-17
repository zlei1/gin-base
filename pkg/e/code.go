package e

var (
	// Common errors
	Ok                  = &E{Code: 0, Message: "OK"}
	InternalServerError = &E{Code: 10001, Message: "Internal server error"}
	ErrBind             = &E{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &E{Code: 10003, Message: "Param error"}
)
