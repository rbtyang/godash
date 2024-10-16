package dasharr

import (
	"reflect"
	"strings"
	"sync"

	"github.com/rbtyang/godash/dashlog"
	"github.com/spf13/cast"
)

/*
Contain @Editor robotyang at 2023

# Contain 利用反射 判断一个 needle值 是否存在于 haystack集合 当中

@Param haystack：待搜索集合，只能是 array/slice/map

@Param needle：被搜索值，是 haystack[0] 类型的值

# Contain use reflection to determine whether a needle value exists in a haystack set

@Param haystack：Set to be searched, haystack can only be array, slice, or map

@Param needle：Searched value, is type haystack[0]
*/
func Contain(haystack any, needle any) bool {
	return inArrayFunc(haystack, func(hayitem any) bool {
		return reflect.DeepEqual(hayitem, needle)
	})
}

/*
inArrayFunc @Editor robotyang at 2023

# inArrayFunc 利用反射 判断一个 needle值 是否存在于 haystack集合 当中

@Param haystack：待搜索集合，只能是 array/slice/map

@Param f：匹配函数

# inArrayFunc use reflection to determine whether a needle value exists in a haystack set

@Param haystack：Set to be searched, haystack can only be array, slice, or map

@Param f：matching function
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
Include @Editor robotyang at 2023

# Include 判断一个 needle值 是否存在于 haystack切片 当中

@Param haystack：待搜索集合，只能是slice (Only can be slice)

@Param needle：被搜索值，是 haystack[0] 类型的值

@Tips comparable：表示go里面 所有内置的 可以使用==或!=来进行比较的类型集合。如 int、uint、float、bool、struct、指针

# Include Determine whether a needle value exists in the haystack slice

@Param haystack：Set to be searched, can only be slice.

@Param needle：Searched value, is type haystack[0]

@Tips comparable：is represents the set of built-in types in go that can be compared using == or != signs. Such as int, uint, float, bool, struct, point.
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
JoinAny @Editor robotyang at 2023

# JoinAny 将任意类型的切片，格式化为字符串

@Param elems：切片值的类型 仅支持数值、字符串 或者两者的混合

@Param separator：分隔符

# JoinAny Format a slice of any type as a string

@Param elems：Slice value types can only be numeric values, strings, or a mixture of both

@Param sep：is separator

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

/*
Chunk @Editor robotyang at 2023

# Chunk 将数组（array）拆分成多个 size 长度的区块，并将这些区块组成一个新数组。 如果array 无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块。

@Param array：待拆分的数组

@Param size：区块大小（size>0）
*/
func Chunk[T any](slice []T, size uint) [][]T {
	if size <= 0 {
		panic("dasharrr.Chunk: size cannot be zero")
	}
	sizeInt := int(size)
	numChunks := (len(slice) + sizeInt - 1) / sizeInt
	result := make([][]T, numChunks)
	for i := 0; i < numChunks; i++ {
		start := i * sizeInt
		end := start + sizeInt
		if end > len(slice) {
			end = len(slice)
		}
		result[i] = slice[start:end]
	}
	return result
}

/*
FilterBy @Editor robotyang at 2023

# FilterBy 根据用户自定义函数，过滤数组元素（性能最佳）

@Param array：待过滤的数组

@Param userFn：用户自定义过滤函数

@Return 过滤后的数组
*/
func FilterBy[T any](list []T, userFn func(T) bool) []T {
	newList := make([]T, 0, len(list))
	for _, item := range list {
		if userFn(item) {
			newList = append(newList, item)
		}
	}
	return newList
}

/*
Deprecated: FilterByWg @Editor robotyang at 2023（性能不佳，仅供学习）

# FilterByWg 根据用户自定义函数，过滤数组元素

@Param array：待过滤的数组

@Param userFn：用户自定义过滤函数

@Return 过滤后的数组
*/
func FilterByWg[T any](array []T, userFn func(T) bool) []T {
	var mu sync.Mutex
	newArr := make([]T, 0, len(array))
	var wg sync.WaitGroup
	wg.Add(len(array))

	// 启动 goroutine 并发处理
	for _, item := range array {
		go func(item T) {
			defer wg.Done()
			if userFn(item) {
				mu.Lock()
				newArr = append(newArr, item)
				mu.Unlock()
			}
		}(item)
	}

	wg.Wait() // 等待所有 goroutine 执行完毕

	return newArr
}

/*
FilterNull @Editor robotyang at 2023

# FilterNull 过滤切片内的空值元素（如 0、0.00、""、"0"、"0.00"、false、nil、空slice、空map，但未支持数组）

@Param list：待过滤的切片

@Return 过滤后的切片

@Reference https://www.php.net/manual/zh/function.array-filter.php
*/
func FilterNull[T any](list []T) []T {
	return FilterBy(list, func(item T) bool {
		v := reflect.ValueOf(item)
		k := v.Kind()
		switch k {
		case reflect.Pointer:
			v2 := v.Elem()
			k2 := v2.Kind()
			switch k2 {
			case reflect.Slice, reflect.Array, reflect.Map:
				return v2.Len() > 0
			}
		case reflect.Slice, reflect.Array, reflect.Map:
			return v.Len() > 0
		}
		return cast.ToFloat64(item) != 0
	})
}

/*
Deprecated: FilterByChan @Editor robotyang at 2023（性能不佳，仅供学习）

# FilterByChan 根据用户自定义函数，过滤数组元素

@Param array：待过滤的数组

@Param userFn：用户自定义过滤函数

@Return 过滤后的数组
*/
func FilterByChan[T any](array []T, userFn func(T) bool, maxGoroutines int) []T {
	var wg sync.WaitGroup
	chs := make([]chan T, maxGoroutines) // 定义多个channel，个数为最大并发goroutine数
	sm := sync.Map{}

	// 从多个channel中接收元素进行处理
	for i := range chs {
		chs[i] = make(chan T)
		wg.Add(1)
		go func(j int, ch chan T) {
			defer wg.Done()
			var localArr []T // 为每个goroutine创建一个局部slice
			for item := range ch {
				if userFn(item) { // 调用用户自定义函数进行处理
					localArr = append(localArr, item) // 成功则将元素写入局部slice
				}
			}
			sm.Store(j, localArr)
		}(i, chs[i])
	}

	for i, item := range array {
		chs[i%maxGoroutines] <- item // 将元素均衡分布到多个channel中
	}

	// 关闭所有channel
	for _, ch := range chs {
		close(ch)
	}

	wg.Wait()

	res := make([]T, 0)
	// 将所有局部slice的结果合并
	sm.Range(func(i any, item any) bool {
		res = append(res, item.([]T)...)
		return true
	})

	return res // 返回筛选结果
}
