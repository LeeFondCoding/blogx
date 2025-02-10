package model

type Comment struct {
	Model
	Content        string     `gorm:"size:256" json:"content"`
	UserID         uint       `json:"UserID"`
	User           User       `gorm:"foreignKey:UserID" json:"-"`
	ArticleID      uint       `json:"articleID"`
	Article        Article    `gorm:"foreignKey:ArticleID" json:"article"`
	ParentID       *uint      `json:"parentID"`
	Parent         *Comment   `gorm:"foreignKey:ParentID" json:"-"`
	SubCommentList []*Comment `gorm:"foreignKey:ParentID" json:"-"`
	RootParentID   *uint      `json:"rootParentID"`
	DiggCount      int        `json:"diggCount"`
}
