package api

import (
	"blogx/api/log_api"
	"blogx/api/site_api"
)

type Api struct {
	Site site_api.SiteApi
	LogApi log_api.LogApi
}

var App = Api{}
