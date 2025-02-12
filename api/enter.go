package api

import "blogx/api/site"

type Api struct {
	Site site.Site
}

var App = Api{}
