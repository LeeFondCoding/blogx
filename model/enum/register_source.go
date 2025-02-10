package enum

type RegisterSourceType int8

const (
	RegisterEmailSourceType    RegisterSourceType = iota + 1
	RegisterQQSourceType        
	RegisterTerminalSourceType  
)