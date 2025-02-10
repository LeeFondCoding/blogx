package main

import (
	"blogx/core"
	"blogx/flag"
	"blogx/global"

	"github.com/sirupsen/logrus"
	//"fmt"
)

func main() {
	flag.Parse()
	global.Conf = core.InitConf()
	core.InitLogrus()

	logrus.Debug("XXX")
	logrus.Warn("XXX")
	logrus.Error("XXX")
	logrus.Info("XXX")
	global.DB = core.InitDB()
}
