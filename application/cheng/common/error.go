package common

const (
	ErrLoginRequired = 401 // 未登录
	ErrSucceed       = "success"   // 成功
	ErrInvalidParam  = "param error"   // 参数错误
	ErrInvalidUser   = "not find user"   // 用户不存在
	ErrWrongPass     = "password error"  // 密码错误
	ErrMysqlQuery    = "mysql error" // MySQL 错误
)
