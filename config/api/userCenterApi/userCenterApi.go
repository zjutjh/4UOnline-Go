package userCenterApi

import "4u-go/config/config"

// UserCenterHost 用户中心地址
var UserCenterHost = config.Config.GetString("user.host")

// UserCenterApi 用户中心接口
type UserCenterApi string

// 用户中心接口
const (
	UCRegWithoutVerify UserCenterApi = "api/activation/notVerify"
	UCReg              UserCenterApi = "api/activation"
	VerifyEmail        UserCenterApi = "api/verify/email"
	ReSendEmail        UserCenterApi = "api/email"
	Auth               UserCenterApi = "api/auth"
	RePass             UserCenterApi = "api/changePwd" // nolint:gosec
	RePassWithoutEmail UserCenterApi = "api/repass"
	DelAccount         UserCenterApi = "api/del"
)
