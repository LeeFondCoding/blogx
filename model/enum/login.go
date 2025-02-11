package enum

type LoginType int8

const (
	UserPwdLoginType LoginType = iota + 1
	QQLoginType
	EmailLoginType
)
