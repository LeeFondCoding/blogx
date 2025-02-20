package router

import (
	"blogx/api"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi

	r.GET("logs", middleware.Admin, app.LogListView)
	r.GET("logs/:id", middleware.Admin, app.LogReadView)
	r.DELETE("logs", middleware.Admin, app.LogRemoveView)
}
