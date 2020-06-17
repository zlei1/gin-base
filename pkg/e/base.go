package e

type E struct {
	Code    int
	Message string
}

func (e *E) Error() string {
	return e.Message
}

func DecodeE(err error) (int, string) {
	if err == nil {
		return Ok.Code, Ok.Message
	}

	switch typed := err.(type) {
	case *E:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
