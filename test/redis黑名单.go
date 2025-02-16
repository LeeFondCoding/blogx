// testdata/7.redis黑名单.go
package main

import (
	"blogx/core"
	"blogx/flag"
	"blogx/global"
	"blogx/service/redis_service/redis_jwt"
	"blogx/util/jwt"
	"fmt"
)

func main4() {

	flag.Parse()
	global.Conf = core.InitConf()
	core.InitLogrus()
	global.Redis = core.InitRedis()

	token, err := jwt.GetToken(jwt.MyClaim{
		UserID: 2,
		Role:   1,
	})
	fmt.Println(token, err)
	redis_jwt.TokenBlack(token, redis_jwt.UserBlackType)
	blk, ok := redis_jwt.HasTokenBlack(token)
	fmt.Println(blk, ok)
}