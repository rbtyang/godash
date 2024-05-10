package dashnet

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Ip2binary @Editor robotyang at 2023

# Ip2binary ip转二进制

@Param ip ip为字符串，如"192.168.56.4"

@Return "11000000101010000011100000000100"
*/
func Ip2binary(ip string) string {
	str := strings.Split(ip, ".")
	var ipstr string
	for _, s := range str {
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			panic("dashnet.Ip2binary: ParseUint: " + err.Error())
		}
		ipstr = ipstr + fmt.Sprintf("%08b", i)
	}
	return ipstr
}

/*
MatchIp @Editor robotyang at 2023

# MatchIp 判断 ip地址 和 其他ip/ip段 是否匹配

@Param ip ip为字符串，如"192.168.56.4"

@Param iprange 为ip段，如"192.168.56.64/26" [64~127, 网络64, 第一65, 最后126, 广播127]

@Param ip ip地址

@Reference https://tool.520101.com/wangluo/ipjisuan/
*/
//
func MatchIp(ip, iprange string) bool {
	ipb := Ip2binary(ip)
	if strings.Contains(iprange, "/") { //如果是ip段
		ipr := strings.Split(iprange, "/")
		masklen, err := strconv.ParseUint(ipr[1], 10, 32)
		if err != nil {
			panic("dashnet.Ip2binary: MatchIp: " + err.Error())
		}
		iprb := Ip2binary(ipr[0])
		return strings.EqualFold(ipb[0:masklen], iprb[0:masklen])
	} else {
		return ip == iprange
	}
}
