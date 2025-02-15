package log_service

import (
	"blogx/core"
	"blogx/global"
	"blogx/model"
	"blogx/model/enum"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType enum.LoginType) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	token := c.GetHeader("token")
	fmt.Println(token)
	// TODO: 通过JWT获取用户ID
	userID := uint(1)
	userName := ""

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
