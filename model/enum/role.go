package enum

// 登录的类型
type Role int8

const (
	AdminType Role = iota + 1
	UserType
	VisitorType
)