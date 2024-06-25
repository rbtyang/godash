package dashstr

import (
	"regexp"
	"strings"
)

type TelType int

const (
	TelType_ZH   TelType = iota + 1 //大陆手机号
	TelType_ZHLD                    //大陆座机号
	TelType_HK                      //香港手机号
	TelType_AM                      //澳门手机号
	TelType_TW                      //台湾手机号
	TelType_QT                      //其他手机号
)

/*
ParseTelType

# ParseTelType 判断手机号类型（是否为港/澳/台/大陆手机or其他）
*/
func ParseTelType(tel string) TelType {
	tel = TrimBlank(tel)
	if IsHKMobile(tel) {
		return TelType_HK
	} else if IsAMMobile(tel) {
		return TelType_AM
	} else if IsTWMobile(tel) {
		return TelType_TW
	} else if IsZHMobile(tel) {
		return TelType_ZH
	} else if IsZHLandline(tel) {
		return TelType_ZHLD
	} else {
		return TelType_QT
	}
}

/*
IsZHMobile

# IsZHMobile 判断是否中国大陆手机号（为+86或86开头或无，可选加-，再以1开头，再加10位数字）

@Tips 建议先使用 TrimBlank 函数过滤特殊字符，再来调用这个方法

@Reference https://mp.weixin.qq.com/s?__biz=MzA3MjMwMzg2Nw==&mid=2247497438&idx=2&sn=2a641c223490cc8073cc7b2e0b940182&chksm=9f22e34aa8556a5c2c81b798daa5a0e784faac088f4d93acc57b2e092704482ae929caf34607&scene=27
*/
func IsZHMobile(tel string) bool {
	reg := `^(?:\+?86\-?)?1\d{10}$`
	bol, _ := regexp.MatchString(reg, tel)
	return bol
}

/*
IsZHLandline

# IsZHLandline 判断是否中国大陆座机号（为0开头的3位或4位区号，可选加-，再加7位或8位电话号）

@Tips 建议先使用 TrimBlank 函数过滤特殊字符，再来调用这个方法
*/
func IsZHLandline(tel string) bool {
	if IsHATMobile(tel) { //港澳台
		return false
	}

	p1 := `^0\d{2,3}\-?`
	re1 := regexp.MustCompile(p1)
	str1 := re1.FindString(tel)

	p2 := `^\d{7,8}$`
	str2 := strings.TrimPrefix(tel, str1)
	bol2, _ := regexp.MatchString(p2, str2)

	return bol2
}

/*
IsHATMobile

# IsHATMobile 是否港澳台号码
*/
func IsHATMobile(tel string) bool {
	return IsHKMobile(tel) || IsAMMobile(tel) || IsTWMobile(tel)
}

/*
IsHKMobile

# IsHKMobile 判断是否中国香港号码（为区号00852开头，可选加-，再加8位数字）

@Tips 建议先使用 TrimBlank 函数过滤特殊字符，再来调用这个方法
*/
func IsHKMobile(tel string) bool {
	return strings.HasPrefix(tel, "00852")
}

/*
IsAMMobile

# IsAMMobile 判断是否中国澳门号码

@Tips 建议先使用 TrimBlank 函数过滤特殊字符，再来调用这个方法
*/
func IsAMMobile(tel string) bool {
	return strings.HasPrefix(tel, "00853")
}

/*
IsTWMobile

# IsTWMobile 判断是否中国台湾号码

@Tips 建议先使用 TrimBlank 函数过滤特殊字符，再来调用这个方法
*/
func IsTWMobile(tel string) bool {
	return strings.HasPrefix(tel, "00886")
}
