package router

import (
	"blogx/global"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	nr := r.Group("api")
	SiteRouter(nr)

	nr.Use(middleware.Log)

	addr := global.Conf.System.Addr()
	r.Run(addr)
}
