package model

type Log struct {
	Model
	LogType int8   `json:"logType"` // 日志类型 1 2 3
	Title   string `gorm:"size:64" json:"title"`
	Content string `json:"content"`
	Level   int8   `json:"level"`                      // 日志级别 1 2 3
	UserID  uint   `json:"userID"`                     // 用户id
	User    User   `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP      string `gorm:"size:32" json:"ip"`
	Addr    string `gorm:"size:64" json:"addr"`
	IsRead  bool   `json:"isRead"` // 是否读取
}
