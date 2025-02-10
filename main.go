package main

import (
	"blogx/core"
	"blogx/flag"
	"blogx/global"
)

func main() {
	flag.Parse()
	global.Conf = core.InitConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	flag.Run()
}
