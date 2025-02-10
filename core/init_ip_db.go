package core

import (
	ipUtil "blogx/util/ip"
	_ "embed"
	"fmt"
	"strings"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

var searcher *xdb.Searcher

//go:embed ip2region.xdb
var addrDB []byte

func InitIPDB() {
	_searcher, err := xdb.NewWithBuffer(addrDB)
	if err != nil {
		logrus.Fatalf("ip地址数据库加载失败 %s", err)
	}
	searcher = _searcher
}

func GetIpAddr(ip string) (addr string) {
	if ipUtil.IsLocalIPAddr(ip) {
		return "内网"
	}

	region, err := searcher.SearchByStr(ip)
	if err != nil {
		logrus.Warnf("异常的IP地址 %s", err)
	}
	// addrlist: 国家|0|省份|市|运营商
	addrList := strings.Split(region, "|")
	//fmt.Println(addrList)

	country, province, city := addrList[0], addrList[2], addrList[3]
	if province != "0" && city != "0" {
		return fmt.Sprintf("%s·%s", province, city)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		return country
	}
	//return region
	return "未知"
}
