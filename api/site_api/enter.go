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
		rep := global.Conf.Email
		rep.AuthCode = "******"
		data = rep
	case "qq":
		rep := global.Conf.QQ
		rep.AppKey = "******"
		data = rep
	case "qiNiu":
		rep := global.Conf.QiNiu
		rep.SecretKey = "******"
		data = rep
	case "ai":
		rep := global.Conf.Ai
		rep.SecretKey = "******"
		data = rep
	}
	res.OkWithData(data, c)
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required" label:"年龄"`
}

func (SiteApi) SiteInfoQQView(c *gin.Context) {
	res.OkWithData(global.Conf.QQ.Url(), c)
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
