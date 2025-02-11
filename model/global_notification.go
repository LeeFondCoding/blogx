package model

type GlobalNotification struct {
	Model
	Title   string `gorm:"size:32" json:"title"`
	Icon    string `gorm:"size:256" json:"icon"`
	Content string `gorm:"size:64" json:"content"`
	Href    string `gorm:"size:256" json:"href"` // 用户点击消息，然后进行跳转
}
