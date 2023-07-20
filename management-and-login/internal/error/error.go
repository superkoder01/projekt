package error

type Error struct {
	Code    int
	Message string
	Cause   error
	Type    string
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Message
}

func New(msg string, code int, typ ErrorType) error {
	return &Error{
		Message: msg,
		Code:    code,
		Type:    typ.Name(),
	}
}

func Wrap(e error, code int) error {
	return &Error{
		Cause: e,
		Code:  code,
	}
}
