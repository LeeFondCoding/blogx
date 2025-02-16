package redis_jwt

import (
	"blogx/global"
	"blogx/util/jwt"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type BlackType int8

const (
	UserBlackType   = iota + 1 // 用户注销
	AdminBlackType             // 管理员手动下线
	DeviceBlackType            // 其他设备挤下线
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}

func ParseBlackType(val string) BlackType {
	switch val {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}

func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)

	claim, err := jwt.ParseToken(token)
	if err != nil || claim == nil {
		logrus.Errorf("token解析失败 %s", err)
		return
	}
	second := claim.ExpiresAt - time.Now().Unix()

	_, err = global.Redis.Set(key, value.String(), time.Duration(second)*time.Second).Result()
	if err != nil {
		logrus.Errorf("redis添加黑名单失败 %s", err)
		return
	}
}

func HasTokenBlack(token string) (blk BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	blk = ParseBlackType(value)
	return blk, true
}
