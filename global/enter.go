package global

import (
	"blogx/conf"

	"gorm.io/gorm"
)

var (
	Conf *conf.Config
	DB   *gorm.DB
)
