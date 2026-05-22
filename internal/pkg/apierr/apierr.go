package apierr

type CodeError struct {
	Code int
	Msg  string
}

func (e *CodeError) Error() string {
	return e.Msg
}

func New(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}
