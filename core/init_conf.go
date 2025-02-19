package core

import (
	"blogx/conf"
	"blogx/flag"
	"blogx/global"

	"fmt"
	"os"

	"github.com/sirupsen/logrus"
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

// 重新格式化配置文件
func SetConf() {
	byteData, err := yaml.Marshal(global.Conf)
	if err != nil {
		logrus.Errorf("conf读取失败 %s", err)
		return
	}

	err = os.WriteFile(flag.Option.File, byteData, 0666)
	if err != nil {
		logrus.Errorf("设置配置文件失败 %s", err)
		return
	}
}
