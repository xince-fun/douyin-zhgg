package errno

// 规定一下错误码
/**(前4位代表业务,后1位代表具体功能)**/
// USER相关错误 1001*开头 包括Token等等
// FEED相关错误 1002*开头
// COMMENT相关错误 1003*开头
// 其他错误以此类推
// 全局错误码
const (
	SuccessCode    = 0
	ServiceErrCode = 10001
	ParamErrCode   = 10002
)

// 用户模块
const (
	UserAlreadyExistErrCode    = 10010
	UserNotExistErrCode        = 10011
	AuthorizationFailedErrCode = 10012
	TokenInvalidErrCode        = 10013
	LoginFailErrCode           = 10014
)

// 视频模块
const (
	PublishVideoToPublicErrCode = 10020
	PublishVideoToOssErrCode    = 10021
	PublishCoverToOssErrCode    = 10022
	GetVideoCoverErrCode        = 10023
	GetOssVideoUrlErrCode       = 10024
	GetOssCoverUrlErrCode       = 10025
)

//
