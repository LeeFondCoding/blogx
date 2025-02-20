package global

import (
	"blogx/conf"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
)

const Version = "1.0.0"

var (
	Conf  *conf.Config
	DB    *gorm.DB
	Redis *redis.Client
	CaptchaStore = base64Captcha.DefaultMemStore
)
