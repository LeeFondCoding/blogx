package site_api

import (
	"blogx/common/res"
	"blogx/global"
	"blogx/middleware"

	"github.com/gin-gonic/gin"
)

type SiteApi struct{}

type SiteInfoRequest struct {
	Name string `uri:"name"`
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if cr.Name == "site" {
		res.OkWithData(global.Conf.Site, c)
		return
	}

	// 判断角色是不是管理员
	middleware.Admin(c)
	_, ok := c.Get("claims")
	if !ok {
		return
	}

	var data any
	switch cr.Name {
	case "email":
		data = global.Conf.Email
	case "qq":
		data = global.Conf.QQ
	case "qiNiu":
		data = global.Conf.QiNiu
	case "ai":
		data = global.Conf.Ai
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	res.OkWithData(data, c)
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required" label:"年龄"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {

	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	res.OkWithMsg("更新成功", c)
}
