package router

import (
	"blogx/api"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi

	r.GET("banner", app.BannerListView)
	r.POST("banner", middleware.Admin, app.BannerCreateView)
	r.PUT("banner/:id", middleware.Admin, app.BannerUpdateView)
	r.DELETE("banner", middleware.Admin, app.BannerRemoveView)
}
