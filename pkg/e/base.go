package e

type E struct {
	Code    int
	Message string
}

func (e *E) Error() string {
	return e.Message
}
