package api

import (
	"blogx/api/banner_api"
	"blogx/api/captcha_api"
	"blogx/api/image_api"
	"blogx/api/log_api"
	"blogx/api/site_api"
	"blogx/api/user_api"
)

type Api struct {
	Site site_api.SiteApi
	LogApi log_api.LogApi
	ImageApi image_api.ImageApi
	BannerApi banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
	User user_api.UserApi
}

var App = Api{}
