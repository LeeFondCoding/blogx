package router

import (
	"blogx/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.User
	r.POST("user/send_email", app.SendEmailView)
}