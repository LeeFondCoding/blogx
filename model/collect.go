package model

// 收藏表
type Collect struct {
	Model
	Title        string `gorm:"size:32" json:"title"`
	Abstract     string `gorm:"size:256" json:"abstract"`
	Cover        string `gorm:"size:256" json:"cover"`
	ArticleCount int    `json:"articleCount"`
	UserID       uint   `json:"userID"`
	User         User   `gorm:"foreignKey:UserID" json:"-"`
}
