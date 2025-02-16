package enum

// 日志的类型：登录日志，运行日志
type LogType int8

const (
	LoginLogType LogType = iota + 1
	ActionLogType
	RuntimeLogType
)
