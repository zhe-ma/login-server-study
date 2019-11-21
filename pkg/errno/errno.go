package errno

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Detail  error
}

func New(errno *Errno, detail error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Detail:  detail,
	}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, detail: %s",
		err.Code, err.Message, err.Detail)
}

func DecodeError(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	}

	return InternalServerError.Code, err.Error()
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeError(err)
	return code == ErrUserNotFound.Code
}
