package site

import (
	"blogx/model/enum"
	"blogx/service/log"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Site struct{}

func (Site) SiteInfoView(c *gin.Context) {
	fmt.Println("1")
	log.NewLoginSuccess(c, enum.UserPwdLoginType)
	log.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "lichun", "1234")
	c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
}
