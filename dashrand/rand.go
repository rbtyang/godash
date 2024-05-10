package dashrand

import (
	"math/rand"
	"reflect"
	"time"
)

// RandMode 组合模式
type RandMode string

const (
	ModeNum        RandMode = "0123456789"                       //整数
	Modeaz         RandMode = "abcdefghijklmnopqrstuvwxyz"       //小写字母
	ModeAZ         RandMode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"       //大写字母
	ModeSp         RandMode = "~!@#$%^&*()_+"                    //特殊字符
	ModeNumAlpha            = ModeNum + Modeaz + ModeAZ          //整数+小写字母+大写字母
	ModeNumAlphaSp          = ModeNum + Modeaz + ModeAZ + ModeSp //整数+小写字母+大写字母+特殊字符
)

/*
Str @Editor robotyang at 2023

# Str 生成随机字符串

@Param mode 字符组合模式

@Param length 要生成的字符长度

@Return str 随机字符
*/
func Str(mode RandMode, length uint) string {
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())

	var charLib = []rune(mode)
	charLibLen := len(charLib)

	randRune := make([]rune, length)
	for i := range randRune {
		randRune[i] = charLib[rand.Intn(charLibLen)]
	}
	return string(randRune)
}

/*
Num @Editor robotyang at 2023

# Num 生成范围内的 随机数

@Param min 下限（包含）

@Param max 上限（不包含）

@Return [min, max)
*/
func Num[T int | int64 | float64](min, max T) T {
	switch reflect.TypeOf(min).Kind() {
	case reflect.Int:
		n := rand.Intn(int(max)-int(min)) + int(min)
		return T(n)
	case reflect.Int64:
		n := rand.Int63n(int64(max)-int64(min)) + int64(min)
		return T(n)
	case reflect.Float64:
		n := rand.Float64()*(float64(max)-float64(min)) + float64(min)
		return T(n)
	default:
		panic("dashrand.Num: unsupported type")
	}
}

/*
NumSlice @Editor robotyang at 2023

# NumSlice 生成范围内的 随机数切片

@Param min 下限（包含）

@Param max 上限（不包含）
*/
func NumSlice[T int | int64 | float64](length uint, min, max T) []T {
	sli := make([]T, length)
	for i := uint(0); i < length; i++ {
		sli[i] = Num(min, max)
	}
	return sli
}
