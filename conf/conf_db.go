package conf

import (
	"fmt"
)

type DB struct {
	User     string `yaml:"user"`
	PassWord string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBname   string `yaml:"dbname"`
	Debug    bool   `yaml:"debug"`
}

func (db *DB) DSN() string {
	//fmt.Printf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local\n", db.User, db.PassWord, db.IP, db.Port, db.DBname)
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.User, db.PassWord, db.Host, db.Port, db.DBname)
}

func (db *DB) Empty() bool {
	return db.User == "" && db.PassWord == "" && db.Host == "" && db.Port == 0
}
