package log_service

import (
	"blogx/core"
	"blogx/global"
	"blogx/model"
	"blogx/model/enum"
	"blogx/util/jwt"

	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	claim, err := jwt.ParseTokenByGin(c)
	var userID uint
	var userName string
	if err == nil && claim != nil {
		userID = claim.UserID
		userName = claim.Username
	}

	global.DB.Create(&model.Log{
		LogType:     enum.LoginLogType,
		Title:       "用户登录",
		UserID:      userID,
		IP:          ip,
		Addr:        addr,
		LoginStatus: true,
		UserName:    userName,
		LoginType:   loginType,
	})
}

func NewLoginFail(c *gin.Context, loginType enum.LoginType, msg, username, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	global.DB.Create(&model.Log{
		LogType:     enum.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		IP:          ip,
		Addr:        addr,
		LoginStatus: false,
		UserName:    username,
		Pwd:         pwd,
		LoginType:   loginType,
	})
}
