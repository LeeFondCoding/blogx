package flag

import (
	"flag"
	"os"
)

type Options struct {
	File    string
	DB      bool
	Version bool
}

var Option = new(Options)

func Parse() {
	flag.StringVar(&Option.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&Option.DB, "db", false, "数据库迁移")
	flag.BoolVar(&Option.Version, "v", false, "版本")
	flag.Parse()
}

func Run() {
	if Option.DB {
		FlagDB()
		os.Exit(0)
	}
}
