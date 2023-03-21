package dasharr

import (
	"github.com/rbtyang/godash/dashlog"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

/*
@Editor robotyang at 2023

Contain 利用反射 判断一个 needle值 是否存在于 haystack集合 当中

@Param haystack 只能是 array/slice/map

@Param needle 是 haystack[0] 类型的值

Contain use reflection to determine whether a needle value exists in a haystack set

@Param haystack can only be array, slice, or map

@Param needle is type haystack[0]
*/
func Contain(haystack any, needle any) bool {
	return inArrayFunc(haystack, func(hayitem any) bool {
		return reflect.DeepEqual(hayitem, needle)
	})
}

/*
@Editor robotyang at 2023

inArrayFunc is a ...
*/
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
		dashlog.Panicf("haystack type must be array/slice/map")
	}
	return false
}

/*
@Editor robotyang at 2023

Include 判断一个 needle值 是否存在于 haystack切片 当中

@Param haystack 只能是slice (Only can be slice)

@Param needle 是 haystack[0] 类型的值

@Tips comparable 表示go里面 所有内置的 可以使用==或!=来进行比较的类型集合。如 int、uint、float、bool、struct、指针

Include Determine whether a needle value exists in the haystack slice

@Param haystack can only be slice.

@Param needle is type haystack[0]

@Tips comparable is represents the set of built-in types in go that can be compared using == or != signs. Such as int, uint, float, bool, struct, point.
*/
func Include[T comparable](haystack []T, needle T) bool {
	for _, ele := range haystack {
		if ele == needle {
			return true
		}
	}
	return false
}

/*
@Editor robotyang at 2023

@Reference strings.Join

JoinAny 将任意类型的切片，格式化为字符串

@Param elems 切片值的类型 仅支持数值、字符串 或者两者的混合

@Param separator 分隔符

JoinAny Format a slice of any type as a string

@Param elems Slice value types can only be numeric values, strings, or a mixture of both

@Param sep is separator
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
