package dashrand

import (
	"math"
	"math/rand"
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

// Num 生成范围内的 随机数

/*
@Editor robotyang at 2023

Num  @return [min, max)
*/
func Num(min, max int) int {
	return rand.Intn(max-min) + min
}

/*
@Editor robotyang at 2023

NumCode  NumCode 生成固定长度的 随机数
*/
func NumCode(len int) int {
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
