package main

import (
	"blogx/core"
	"blogx/flag"
	"blogx/global"
	"blogx/router"
)

func main() {
	flag.Parse()
	global.Conf = core.InitConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	core.InitIPDB()
	global.Redis = core.InitRedis()

	flag.Run()

	router.Run()
}
