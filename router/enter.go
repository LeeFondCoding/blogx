package router

import (
	"blogx/global"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	nr := r.Group("api")
	SiteRouter(nr)

	addr := global.Conf.System.Addr()
	r.Run(addr)
}
