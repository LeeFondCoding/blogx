package user_api

import (
	"blogx/common/res"
	"blogx/global"
	"blogx/model"
	"blogx/service/email_service"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"oneof=1 2"` // 1 注册 2 重置密码
	Email string `json:"email" binding:"required"`
}
type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	code := base64Captcha.RandText(4, "0123456789")
	id := base64Captcha.RandomId()

	switch cr.Type {
	case 1:
		// 查邮箱是否不存在
		var user model.User
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("该邮箱以使用", c)
			return
		}
		err = email_service.SendRegisterCode(cr.Email, code)
	case 2:
		err = email_service.SendResetPwdCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败 %s", err)
		res.FailWithMsg("邮件发送失败", c)
		return
	}
	global.CaptchaStore.Set(id, code)
	res.OkWithData(SendEmailResponse{
		EmailID: id,
	}, c)
}