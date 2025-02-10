package model

type Image struct {
	Model
	FileName string `gorm:"size:64" json:"filename"`
	Path     string `gorm:"size:256" json:"path"`
	Size     int64  `json:"size"`
	Hash     string `gorm:"size:32" json:"hash"`
}

func (i Image) WebPath() string {
	return "/"
}
