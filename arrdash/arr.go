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
SliceHas 判断一个 needle值 是否存在于 haystack切片 当中

@Param haystack 只能是 slice

@Param needle 是 haystack[0] 类型的值

@Tips comparable 表示go里面 所有内置的 可比较类型：int、uint、float、bool、struct、指针 等一切可以比较的类型
*/
func SliceHas[T comparable](haystack []T, needle T) bool {
	for _, ele := range haystack {
		if ele == needle {
			return true
		}
	}
	return false
}

/*
Contain 利用反射 判断一个 needle值 是否存在于 haystack集合 当中

@Param haystack 只能是 slice/array/map

@Param needle 是 haystack[0] 类型的值
*/
func Contain(haystack any, needle any) bool {
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
		logdash.Errorf("arrdash.Contain haystack type must be slice/array/map, Yours %#v", haystack)
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
