package router

import (
	"blogx/api"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Site
	r.GET("site/qq_url", app.SiteInfoQQView)
	r.GET("site/:name", app.SiteInfoView)
	r.PUT("site/:name", middleware.Admin, app.SiteUpdateView)
}
