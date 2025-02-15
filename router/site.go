package router

import (
	"blogx/api"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Site
	r.GET("site", app.SiteInfoView)
	r.PUT("site", app.SiteUpdateView)
}
