package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func (e ErrNo) GetErrCode() int64 {
	return e.ErrCode
}

func (e ErrNo) GetErrMsg() string {
	return e.ErrMsg
}

var (
	Success    = NewErrNo(SuccessCode, "Success")
	ServiceErr = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr   = NewErrNo(ParamErrCode, "Wrong Parameter has been given")

	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	UserNotExistErr        = NewErrNo(UserNotExistErrCode, "User does not exist")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	TokenInvalidErr        = NewErrNo(TokenInvalidErrCode, "Token invalid")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
