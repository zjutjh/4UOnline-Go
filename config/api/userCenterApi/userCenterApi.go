package userCenterApi

import "4u-go/config/config"

// UserCenterHost 用户中心地址
var UserCenterHost = config.Config.GetString("user.host")

// 用户中心接口
const (
	UCRegWithoutVerify string = "api/activation/notVerify"
	UCReg              string = "api/activation"
	VerifyEmail        string = "api/verify/email"
	ReSendEmail        string = "api/email"
	Auth               string = "api/auth"
	RePass             string = "api/changePwd" // nolint:gosec
	RePassWithoutEmail string = "api/repass"
	DelAccount         string = "api/del"
)
