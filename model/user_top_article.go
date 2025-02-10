package model

import (
	"time"
)

type UserTopArticle struct {
	UserID    uint      `gorm:"uniqueIndex:idx_name" json:"userID"`
	ArticleID uint      `gorm:"uniqueIndex:idx_name" json:"articleID"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Article   Article   `gorm:"foreignKey:ArticleID" json:"-"`
	CreatedAt time.Time `json:"createAt"`
}
