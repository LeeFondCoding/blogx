package enum

// 哪种方式登录的
type LoginType int8

const (
	UserPwdLoginType LoginType = iota + 1
	QQLoginType
	EmailLoginType
)
