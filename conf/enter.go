package conf

type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
	Jwt    Jwt    `yaml:"jwt"`
	DB     DB     `yaml:"db"`  // 读库
	DB1    DB     `yaml:"db1"` // 写库
	Redis  Redis  `yaml:"redis"`
}
