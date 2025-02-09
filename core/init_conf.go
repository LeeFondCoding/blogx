package core

import (
	"blogx/flag"
	
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

var configPath = "settings.yaml"
type System struct {
	Ip string `yaml:"IP"`
	Port int `yaml:"Port"`
}

type Config struct {
	System System `yaml:"system"`
}

func InitConf() {
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(data))
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Sprintf("yaml配置文件给是错误 %s", err))
	}

	fmt.Printf("成功读取配置文件: %s", flag.Option.File)
}