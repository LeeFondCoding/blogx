package router

import (
	"blogx/api"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Site
	r.GET("site", app.SiteInfoView)
	r.PUT("site", middleware.Admin, app.SiteUpdateView)
}
