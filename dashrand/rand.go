package dashrand

import (
	"math"
	"math/rand"
	"reflect"
	"time"
)

type RandMode string

const (
	ModeNum        RandMode = "0123456789"                 //数字
	Modeaz         RandMode = "abcdefghijklmnopqrstuvwxyz" //小写字符
	ModeAZ         RandMode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" //大写字符
	ModeSp         RandMode = "!@#$%^&*()_+~"              //特殊字符
	ModeNumAlpha            = ModeNum + Modeaz + ModeAZ
	ModeNumAlphaSp          = ModeNum + Modeaz + ModeAZ + ModeSp
)

/*
@Editor robotyang at 2023

Str  @return str
*/
func Str(mode RandMode, n uint16) string {
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	randRune := make([]rune, n)
	var chList = []rune(mode)
	length := len(chList)
	for i := range randRune {
		randRune[i] = chList[rand.Intn(length)]
	}
	return string(randRune)
}

/*
@Editor robotyang at 2023

Num 生成范围内的 随机数

@return [min, max)
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
@Editor robotyang at 2023

NumLen 生成固定长度的 随机数
*/
func NumLen(len int) int {
	min := math.Pow10(len - 1)
	max := min * 10
	return Num(int(min), int(max))
}

/*
@Editor robotyang at 2023

IntSli is a ...
*/
func IntSli(len, min, max int) []int {
	sli := make([]int, len)
	for i := 0; i < len; i++ {
		sli[i] = Num(min, max)
	}
	return sli
}
