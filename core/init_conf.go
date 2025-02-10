package core

import (
	"blogx/conf"
	"blogx/flag"

	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var configPath = "settings.yaml"

func InitConf() (c *conf.Config) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(data))
	c = new(conf.Config)
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件给是错误 %s", err))
	}

	fmt.Printf("成功读取配置文件: %s\n", flag.Option.File)
	return
}
