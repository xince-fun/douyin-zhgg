package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func (e ErrNo) GetErrCode() int32 {
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

	PublishVideoToPublicErr = NewErrNo(PublishVideoToPublicErrCode, "Unable to public video to public")
	PublishVideoToOssErr    = NewErrNo(PublishVideoToOssErrCode, "Unable to public video to oss")
	PublishCoverToOssErr    = NewErrNo(PublishCoverToOssErrCode, "Unable to public cover to oss")
	GetVideoCoverErr        = NewErrNo(GetVideoCoverErrCode, "Unable to get video cover")
	GetOssVideoUrlErr       = NewErrNo(GetOssVideoUrlErrCode, "Unable to get video url from oss")
	GetOssCoverUrlErr       = NewErrNo(GetOssCoverUrlErrCode, "Unable to get video cover url from oss")
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
