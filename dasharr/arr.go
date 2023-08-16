package dasharr

import (
	"github.com/rbtyang/godash/dashlog"
	"github.com/spf13/cast"
	"reflect"
	"strings"
	"sync"
)

/*
# Chunk 将切片拆分成多个 size 长度的区块，并将这些区块组成一个新数组。如果切片无法被分割成全部等长的区块，那么最后剩余的元素将组成一个区块。

@Param slice 等待拆分的切片

@Param size 区块大小（size>0）

# Chunk Split the slice into multiple size-length blocks and compose the blocks into a new array.If the slice cannot be divided into all equal length blocks, then the last remaining elements will form a block.

@Param slice Slices waiting to be split

@Param size Block size (size>0)
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
# Contain 利用反射 判断一个 needle值 是否存在于 haystack集合 当中

@Param haystack 只能是 array/slice/map

@Param needle 是 haystack[0] 类型的值

# Contain use reflection to determine whether a needle value exists in a haystack set

@Param haystack can only be array, slice, or map

@Param needle is type haystack[0]
*/
func Contain(haystack any, needle any) bool {
	return inArrayFunc(haystack, func(hayitem any) bool {
		return reflect.DeepEqual(hayitem, needle)
	})
}

/*
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
FilterBy 根据用户自定义函数，过滤数组元素（性能最佳）

@Param array 待过滤的数组

@Param userFn 用户自定义过滤函数

@Return 过滤后的数组
*/
func FilterBy[T any](array []T, userFn func(T) bool) []T {
	newArr := make([]T, 0, len(array))
	for _, item := range array {
		if userFn(item) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

/*
Deprecated: FilterByWg 根据用户自定义函数，过滤数组元素（性能不佳，仅供学习）

@Param array 待过滤的数组

@Param userFn 用户自定义过滤函数

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

func FilterNull[T any](array []T) []T {
	return FilterByWg(array, func(item T) bool {
		return cast.ToInt(item) == 0
	})
}

/*
Deprecated: FilterByChan 根据用户自定义函数，过滤数组元素（性能不佳，仅供学习）

@Param array 待过滤的数组

@Param userFn 用户自定义过滤函数

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

	var res []T
	wg.Wait()

	// 将所有局部slice的结果合并
	sm.Range(func(i any, item any) bool {
		res = append(res, item.([]T)...)
		return true
	})

	return res // 返回筛选结果
}

/*
# Include 判断一个 needle值 是否存在于 haystack切片 当中

@Param haystack 只能是slice (Only can be slice)

@Param needle 是 haystack[0] 类型的值

@Tips comparable 表示go里面 所有内置的 可以使用==或!=来进行比较的类型集合。如 int、uint、float、bool、struct、指针

# Include Determine whether a needle value exists in the haystack slice

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
# JoinAny 将任意类型的切片，格式化为字符串

@Param elems 切片值的类型 仅支持数值、字符串 或者两者的混合

@Param separator 分隔符

# JoinAny Format a slice of any type as a string

@Param elems Slice value types can only be numeric values, strings, or a mixture of both

@Param sep is separator

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
