package dashstr

import (
	"regexp"
	"strings"
	"unicode"
)

/*
UpperFirst  首字母大写
*/
func UpperFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(string([]rune(s)[0])) + string([]rune(s)[1:])
}

/*
驼峰转蛇形 snake string
@description XxYy to xx_yy , XxYY to xx_y_y
@param s 需要转换的字符串
@Return string
*/
func SnakeString(s string) string {
	return snakeStr(s, '_')
}

/*
驼峰转蛇形 snake string
@description XxYy to xx-yy , XxYY to xx-y-y
@param s 需要转换的字符串
@Return string
*/
func SnakeStringStrike(s string) string {
	return snakeStr(s, '-')
}

/*
snakeStr  通用转换
*/
func snakeStr(s string, sep byte) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, sep)
		}
		if d != sep {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

/*
蛇形转驼峰
@description xx_yy to XxYx  xx_y_y to XxYY
@param s要转换的字符串
@Return string
*/
func CamelStringStrike(s string) string {
	return camelStr(s, '-')
}

/*
蛇形转驼峰
@description xx-yy to XxYx  xx-y-y to XxYY
@param s要转换的字符串
@Return string
*/
func CamelString(s string) string {
	return camelStr(s, '_')
}

/*
camelStr  通用驼峰转换
*/
func camelStr(s string, sep byte) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == sep && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

/*
IsDigit 判断是否为 纯整数

@Param str 待判断的字符
*/
func IsDigit(str string) bool {
	if "" == str {
		return false
	}
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

/*
IsLetter 判断是否为 纯字母

@Param str 待判断的字符
*/
func IsLetter(str string) bool {
	b, err := regexp.MatchString("^([A-z]+)$", str)
	if err != nil {
		panic("dashstr.IsLetter: " + err.Error())
	}
	return b
}

/*
IsDigitLetter 判断是否为 整数+字母

@Param str 待判断的字符
*/
func IsDigitLetter(str string) bool {
	b, err := regexp.MatchString("^([0-9A-z]+)$", str)
	if err != nil {
		panic("dashstr.IsDigitLetter: " + err.Error())
	}
	return b
}

/*
IsCapitalLetter 判断是否为 纯大写字母

@Param str 待判断的字符
*/
func IsCapitalLetter(str string) bool {
	b, err := regexp.MatchString("^([A-Z]+)$", str)
	if err != nil {
		panic("dashstr.IsCapitalLetter: " + err.Error())
	}
	return b
}

/*
IsSmallLetter 判断是否为 纯小写字母

@Param str 待判断的字符
*/
func IsSmallLetter(str string) bool {
	b, err := regexp.MatchString("^([a-z]+)$", str)
	if err != nil {
		panic("dashstr.IsSmallLetter: " + err.Error())
	}
	return b
}
