package model

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Image struct {
	Model
	FileName string `gorm:"size:64" json:"filename"`
	Path     string `gorm:"size:256" json:"path"`
	Size     int64  `json:"size"`
	Hash     string `gorm:"size:32" json:"hash"`
}

func (i Image) WebPath() string {
	return fmt.Sprintf("/%s", i.Path)
}

func (l *Image) BeforeDelete(tx *gorm.DB) error {
	err := os.Remove(l.Path)
	if err != nil {
		logrus.Warnf("删除文件失败 %s", err)
	}
	return nil
}
