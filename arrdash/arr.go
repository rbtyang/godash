package errdash

import (
	"fmt"
	"github.com/rbtyang/godash/logdash"
	"reflect"
	"strings"
)

// Contains 判断一个 needle值 是否存在于 haystack值数组 当中
// @param.haystack must be slice/array/map
func Contains(haystack interface{}, needle interface{}) bool {
	return inArrayFunc(haystack, func(hayitem interface{}) bool {
		return reflect.DeepEqual(hayitem, needle)
	})
}

func inArrayFunc(haystack interface{}, f func(interface{}) bool) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if f(val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if f(val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		logdash.Errorf("arrdash.Contains haystack type must be slice/array/map, Yours %#v", haystack)
	}
	return false
}

// IntsHas check the []int contains the given value
func IntsHas(ints []int, val int) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// Int64sHas check the []int64 contains the given value
func Int64sHas(ints []int64, val int64) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// StringsHas check the []string contains the given element
func StringsHas(ss []string, val string) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
