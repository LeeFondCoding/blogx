// testdata/5.运行日志.go
package main

import (
	"blogx/core"
	"blogx/flag"
	"blogx/global"
	"blogx/service/log_service"
)

func main() {
	flag.Parse()
	global.Conf = core.InitConf()
	core.InitLogrus()
	global.DB = core.InitDB()

	log := log_service.NewRuntimeLog("同步文章数据", log_service.RuntimeDateHour)
	log.SetItem("文章1", 11)
	log.Save()
	log.SetItem("文章2", 12)
	log.Save()
}
