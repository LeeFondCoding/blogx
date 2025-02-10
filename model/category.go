package model

type Category struct {
	Model
	Title  string `gorm:"size:32" json:"title"`
	UserID uint   `json:"userID"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
}
