package site

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Site struct{}

func (Site) SiteInfoView(c *gin.Context) {
    fmt.Println("1")
    c.JSON(200, gin.H{"code": 0, "msg":"站点信息"})
    return
}