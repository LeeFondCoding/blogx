package enum

type LogType int8

const (
	LoginLogType LogType = iota + 1
	ActionLogType
	RuntimeLogType
)
