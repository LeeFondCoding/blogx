package model

type Login struct {
	Model
	UserID uint   `json:"userID"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
	IP     string `gorm:"size:32" json:"ip"`
	Addr   string `gorm:"size:64" json:"addr"`
	UA     string `gorm:"size:128" json:"ua"`
}
