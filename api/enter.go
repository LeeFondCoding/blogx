package api

import "blogx/api/site"

type Api struct {
	Site site.SiteApi
}

var App = Api{}
