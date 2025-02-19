package site_api

import (
	"blogx/common/res"
	"blogx/conf"
	"blogx/core"
	"blogx/global"
	"blogx/middleware"
	"errors"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		global.Conf.Site.About.Version = global.Version
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
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var rep any
	switch cr.Name {
	case "site":
		var data conf.Site
		err = c.ShouldBindJSON(&data)
		rep = data
	case "email":
		var data conf.Email
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qq":
		var data conf.QQ
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qiNiu":
		var data conf.QiNiu
		err = c.ShouldBindJSON(&data)
		rep = data
	case "ai":
		var data conf.Ai
		err = c.ShouldBindJSON(&data)
		rep = data
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	switch s := rep.(type) {
	case conf.Site:
		// 判断站点信息更新前端文件部分
		err = UpdateSite(s)
		if err != nil {
			res.FailWithError(err, c)
			return
		}
		global.Conf.Site = s
	case conf.Email:
		if s.AuthCode == "******" {
			s.AuthCode = global.Conf.Email.AuthCode
		}
		global.Conf.Email = s
	case conf.QQ:
		if s.AppKey == "******" {
			s.AppKey = global.Conf.QQ.AppKey
		}
		global.Conf.QQ = s
	case conf.QiNiu:
		if s.SecretKey == "******" {
			s.SecretKey = global.Conf.QiNiu.SecretKey
		}
		global.Conf.QiNiu = s
	case conf.Ai:
		if s.SecretKey == "******" {
			s.SecretKey = global.Conf.Ai.SecretKey
		}
		global.Conf.Ai = s
	}

	// 改配置文件
	core.SetConf()

	res.OkWithMsg("更新站点配置成功", c)
}

func UpdateSite(site conf.Site) error {
	if site.Project.Icon == "" && site.Project.Title == "" &&
		site.Seo.Keywords == "" && site.Seo.Description == "" &&
		site.Project.WebPath == "" {
		return nil
	}

	if site.Project.WebPath == "" {
		return errors.New("请配置前端地址")
	}

	file, err := os.Open(site.Project.WebPath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s 文件不存在", site.Project.WebPath))
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		logrus.Errorf("goquery 解析失败 %s", err)
		return errors.New("文件解析失败")
	}

	if site.Project.Title != "" {
		doc.Find("title").SetText(site.Project.Title)
	}
	if site.Project.Icon != "" {
		if doc.Is("link[ref='icon']") {
			// 有就修改
			doc.Find("link[ref='icon']").SetAttr("href", site.Project.Icon)
		} else {
			// 没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"%s\">", site.Project.Icon))
		}
	}
	if site.Seo.Keywords != "" {
		if doc.Is("meta[name='keywords']") {
			doc.Find("meta[name='keywords']").SetAttr("content", site.Seo.Keywords)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"keywords\" content=\"%s\">", site.Seo.Keywords))
		}
	}
	if site.Seo.Description != "" {
		if doc.Is("meta[name='description']") {
			doc.Find("meta[name='description']").SetAttr("content", site.Seo.Description)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", site.Seo.Description))
		}
	}

	html, err := doc.Html()
	if err != nil {
		logrus.Errorf("生成html失败 %s", err)
		return errors.New("生成html失败")
	}

	err = os.WriteFile(site.Project.WebPath, []byte(html), 0666)
	if err != nil {
		logrus.Errorf("文件写入失败 %s", err)
		return errors.New("文件写入失败")
	}
	return nil
}
