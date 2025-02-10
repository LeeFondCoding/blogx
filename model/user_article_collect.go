package model

import (
	"time"
)

// 用户文章收藏表
type UserArticleCollectModel struct {
	UserID    uint      `gorm:"uniqueIndex:idx_name" json:"userID"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	ArticleID uint      `gorm:"uniqueIndex:idx_name" json:"articleID"`
	Article   Article   `gorm:"foreignKey:ArticleID" json:"-"`
	CollectID uint      `gorm:"uniqueIndex:idx_name" json:"collectID"`     // 收藏夹的id
	Collect   Collect   `gormf:"foreignKey:CollectID" json:"collectModel"` // 属于哪一个收藏夹
	CreatedAt time.Time `json:"createdAt"`                                 // 收藏的时间
}
