package api

import (
	"blogx/api/image_api"
	"blogx/api/log_api"
	"blogx/api/site_api"
)

type Api struct {
	Site site_api.SiteApi
	LogApi log_api.LogApi
	ImageApi image_api.ImageApi
}

var App = Api{}
