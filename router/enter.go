package router

import (
	"blogx/global"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Conf.System.GinMode)
	r := gin.Default()

	r.Static("/upload", "upload")

	nr := r.Group("api")
	SiteRouter(nr)
	LogRouter(nr)
	ImageRouter(nr)

	nr.Use(middleware.Log)

	addr := global.Conf.System.Addr()
	r.Run(addr)
}
