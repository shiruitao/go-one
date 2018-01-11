package common

const (
	ErrLoginRequired = 401 // 未登录
	ErrSucceed       = 0   // 成功
	ErrInvalidParam  = 1   // 参数错误
	ErrInvalidUser   = 3   // 用户不存在
	ErrWrongPass     = 11  // 密码错误
	ErrMysqlQuery    = 500 // MySQL 错误
)
