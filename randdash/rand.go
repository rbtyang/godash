package randdash

import (
	"math/rand"
	"time"
)

type RandMode string

const (
	ModeNum        RandMode = "0123456789"                 //数字
	Modeaz         RandMode = "abcdefghijklmnopqrstuvwxyz" //小写字符
	ModeAZ         RandMode = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" //大写字符
	ModeSp         RandMode = "!@#$%^&*()_+~"              //特殊字符
	ModeNumAlpha   RandMode = ModeNum + Modeaz + ModeAZ
	ModeNumAlphaSp RandMode = ModeNum + Modeaz + ModeAZ + ModeSp
)

// @return str
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

// @return [min, max)
func Num(min, max int) int {
	return rand.Intn(max-min) + min
}

func IntSli(len, min, max int) []int {
	sli := make([]int, len)
	for i := 0; i < len; i++ {
		sli[i] = Num(min, max)
	}
	return sli
}
