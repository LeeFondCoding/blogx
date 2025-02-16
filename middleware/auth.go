package middleware

import (
	"blogx/common/res"
	"blogx/model/enum"
	"blogx/service/redis_service/redis_jwt"
	"blogx/util/jwt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	claim, err := jwt.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}

	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(),c)
		c.Abort()
		return
	}

	c.Set("claims", claim)
}

func Admin(c *gin.Context) {
	claim, err := jwt.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	if claim.Role != enum.AdminType {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	blcType, ok := redis_jwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}
	c.Set("claims", claim)
}