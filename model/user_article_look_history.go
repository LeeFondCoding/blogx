package model

type UserArticleLookHistory struct {
	Model
	UserID    uint    `json:"userID"`
	User      User    `gorm:"foreignKey:UserID" json:"-"`
	ArticleID uint    `json:"articleID"`
	Article   Article `gorm:"foreignKey:ArticleID" json:"-"`
}
