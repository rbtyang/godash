package dashcrypt

import (
	"bytes"
	"github.com/rbtyang/godash/dashhash"
	"sort"
	"strings"
)

/*
SignFlatMap @Editor robotyang at 2023

# SignFlatMap 根据参数计算签名

@Param data：仅支持 map[string]string / struct / url.Values

@Reference https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Shake_RedPack/Red_Packet_JSAPI.html
*/
func SignFlatMap(data map[string]string, secret string) string {
	// 除sign字段外，所有参数 按照字段名 的ASCII码 从小到大排序后（字典序）
	keys := make([]string, 0, len(data))
	for k := range data {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 使用 URL键值对 的格式（即Param1=value1&Param2=value2…）拼接而成 签名原始串，空值不参与签名组串。
	// 签名原始串中，字段名和字段值 都采用原始值，不进行URL转义。
	var bf bytes.Buffer
	for _, k := range keys {
		v := data[k]
		if v == "" {
			continue
		}
		bf.WriteString(k)
		bf.WriteByte('=')
		bf.WriteString(v)
		bf.WriteByte('&')
	}

	// 签名原始串末尾 补上key参数，如 Param1=value1&Param2=value2…&key=keyvalue
	bf.WriteString("key=")
	bf.WriteString(secret)

	// 得到的字符串 进行MD5，并转换为大写。即 sign=ToUpperCase(MD5(Param1=value1&Param2=value2…&key=keyvalue))
	return strings.ToUpper(dashhash.Md5ByteToStr(bf.Bytes()))
}

/*
CheckSignFlatMap @Editor robotyang at 2023

# CheckSignFlatMap 检查参数签名
*/
func CheckSignFlatMap(data map[string]string, secret, inSign string) bool {
	return inSign == SignFlatMap(data, secret)
}

// //TODO 支持各种结构的签名计算
//func SignObj(data interface{}, secret string) (string, error) {
//	maps, err := convdash.ObjToMap(data)
//	if err != nil {
//		return "", err
//	}
//	switch maps.(type) {
//	case []interface{}:
//		maps := maps.([]interface{})
//		if maps == nil {
//			return "", nil
//		}
//		keys := make([]string, 0, len(maps))
//		if _, ok := maps[0]["SignSort"]; ok {
//			for i, mp := range maps {
//				keys = append(keys, mp["SignSort"])
//			}
//		}
//	case map[string]interface{}:
//		maps := maps.(map[string]interface{})
//		keys := make([]string, 0, len(maps))
//		for k := range maps {
//			if k == "sign" {
//				continue
//			}
//			keys = append(keys, k)
//		}
//		sort.Strings(keys)
//	default:
//		return "", errors.New("Unsupport data type")
//	}
//	return "", nil
//}
//
//// 检查参数签名
//func CheckSignObj(data interface{}, secret, inSign string) bool {
//	newSign, err :=  SignObj(data, secret)
//	if err != nil {
//		return false
//	}
//	return inSign == newSign
//}
