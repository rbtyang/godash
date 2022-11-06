package arrdash

import (
	"fmt"
	"github.com/rbtyang/godash/logdash"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

func init() {
	_ = fmt.Sprintf("fortest")
}

/*
Contains 利用反射 判断一个 needle值 是否存在于 haystack数组当中

@Param haystack 只能是 slice/array/map

@Param needle 是 haystack[0] 类型的值
*/
func Contains(haystack any, needle any) bool {
	return inArrayFunc(haystack, func(hayitem any) bool {
		return reflect.DeepEqual(hayitem, needle)
	})
}

func inArrayFunc(haystack any, f func(any) bool) bool {
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

//IntsHas 检查 []int 包含给定的值
func IntsHas(ints []int, val int) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

//Int64sHas 检查 []int64 包含给定的值
func Int64sHas(ints []int64, val int64) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

//StringsHas 检查 []string 包含给定的值
func StringsHas(strs []string, val string) bool {
	for _, ele := range strs {
		if ele == val {
			return true
		}
	}
	return false
}

/*
JoinAny 将任意类型切片，格式化为字符串（切片值类型 仅支持数值、字符串 或者两者的混合）

@Param separator 分隔符

@Reference strings.Join
*/
func JoinAny(elems []any, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return cast.ToString(elems[0])
	}

	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(cast.ToString(elems[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(cast.ToString(elems[0]))
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(cast.ToString(s))
	}

	return b.String()
}
