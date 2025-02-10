package model

import "blogx/model/enum"

type User struct {
	Model
	UserName       string                  `gorm:"size:32" json:"username"`
	NickName       string                  `gorm:"size:32" json:"nickname"`
	Avatar         string                  `gorm:"size:256" json:"Avastar"` // 头像
	Abstract       string                  `gorm:"size:256" json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"` // 注册来源
	CodeAge        int                     `json:"codeAge"`
	PassWord       string                  `gorm:"size:64" json:"-"`
	Email          string                  `gorm:"size:256" json:"email"` //第三方登录的唯一ID
	OpenID         string                  `gorm:"size:64" json:"openID"`
	Role           enum.Role               `json:"role"` // 角色 1 管理员  2 普通用户  3 访客
}
