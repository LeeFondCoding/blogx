package enum

type Role int8

const (
	AdminType Role = iota + 1
	UserType
	VisitorType
)