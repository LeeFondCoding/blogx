package ip

import "net"

/*
内网地址：
	10.0.0.0 ~ 10.255.255.255
	172.16.0.0 ~ 172.31.255.255
	192.168.0.0 ~ 192.192.255.255
	169.254.0.0 ~ 169.254.255.255
*/

func IsLocalIPAddr(ip string) bool {
	return isLocalIP(net.IP(ip))
}

func isLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
