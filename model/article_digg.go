package model

import "time"

type ArticleDigg struct {
	UserID       uint      `gorm:"uniqueIndex:idx_name" json:"userID"`
	UserModel    User      `gorm:"foreignKey:UserID" json:"-"`
	ArticleID    uint      `gorm:"uniqueIndex:idx_name" json:"articleID"`
	ArticleModel Article   `gorm:"foreignKey:ArticleID" json:"-"`
	CreatedAt    time.Time `json:"createdAt"` // 点赞的时间
}
