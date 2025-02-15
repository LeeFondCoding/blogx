package conf

import "fmt"

type System struct {
	Ip      string `yaml:"ip"`
	Port    int    `yaml:"port"`
	Env     string `yaml:"env"`
	GinMode string `yaml:"gin_mode"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}
