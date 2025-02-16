package main

import (
	"blogx/core"
	"blogx/flag"
	"blogx/global"
	"blogx/util/jwt"
	"fmt"
)

func main() {
	flag.Parse()
	global.Conf = core.InitConf()
	core.InitLogrus()
	token, err := jwt.GetToken(jwt.MyClaim{
		UserID: 2,
		Role:   1,
	})
	fmt.Println(token, err)
}