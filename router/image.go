package router

import (
	"blogx/api"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.Auth, app.ImageUploadView)
	r.GET("images", middleware.Admin, app.ImageListView)
	r.DELETE("images", middleware.Admin, app.ImageRemoveView)
}
