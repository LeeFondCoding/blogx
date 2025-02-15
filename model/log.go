package model

import "blogx/model/enum"

type Log struct {
	Model
	LogType     enum.LogType      `json:"logType"` // 日志类型 1 2 3
	Title       string            `gorm:"size:64" json:"title"`
	Content     string            `json:"content"`
	Level       enum.LogLevelType `json:"level"`                      // 日志级别 1 2 3
	UserID      uint              `json:"userID"`                     // 用户id
	User        User              `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP          string            `gorm:"size:32" json:"ip"`
	Addr        string            `gorm:"size:64" json:"addr"`
	IsRead      bool              `json:"isRead"`                  // 是否读取
	LoginStatus bool              `json:"loginStatus"`             // 登录的状态
	UserName    string            `gorm:"size:32" json:"userName"` // 登录的用户名
	Pwd         string            `gorm:"size:32" json:"username"` // 登录日志的密码
	LoginType   enum.LoginType    `json:"loginType"`               // 登录的类型
}
