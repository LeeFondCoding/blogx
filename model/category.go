package model

type Category struct {
	Model
	Title  string `json:"title"`
	UserID uint   `json:"userID"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
}
