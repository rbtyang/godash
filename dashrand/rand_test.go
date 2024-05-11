package dashrand_test

import (
	"github.com/rbtyang/godash/dashrand"
	"github.com/rbtyang/godash/dashstr"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestStr is a ...
*/
func TestStr(t *testing.T) {
	lenArr := []uint{1, 3, 9, 13, 16, 32}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.ModeNum, leng)
		assert.Equal(t, leng, uint(len(recv)))
		assert.Equal(t, true, dashstr.IsDigit(recv))
		assert.Equal(t, false, dashstr.IsLetter(recv))
	}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.Modeaz, leng)
		assert.Equal(t, leng, uint(len(recv)))
		assert.Equal(t, true, dashstr.IsSmallLetter(recv))
		assert.Equal(t, false, dashstr.IsCapitalLetter(recv))
	}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.ModeAZ, leng)
		assert.Equal(t, leng, uint(len(recv)))
		assert.Equal(t, false, dashstr.IsSmallLetter(recv))
		assert.Equal(t, true, dashstr.IsCapitalLetter(recv))
	}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.ModeNum+dashrand.Modeaz, leng)
		assert.Equal(t, leng, uint(len(recv)))
		assert.Equal(t, true, dashstr.IsDigitLetter(recv))
	}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.ModeNum+dashrand.ModeAZ, leng)
		assert.Equal(t, leng, uint(len(recv)))
		assert.Equal(t, true, dashstr.IsDigitLetter(recv))
	}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.ModeNumAlpha, leng)
		assert.Equal(t, leng, uint(len(recv)))
		assert.Equal(t, true, dashstr.IsDigitLetter(recv))
	}

	for _, leng := range lenArr {
		recv := dashrand.Str(dashrand.ModeNumAlphaSp, leng)
		assert.Equal(t, leng, uint(len(recv)))
		if len(recv) >= 10 {
			assert.Equal(t, false, dashstr.IsDigitLetter(recv))
		}
	}
}

/*
TestNum is a ...
*/
func TestNum(t *testing.T) {
	{
		min, max := -10, 30
		for i := 0; i < 10000; i++ {
			num := dashrand.Num(min, max)
			recv := num >= min && num < max
			if !assert.Equal(t, recv, true) {
				break
			}
		}
	}
	{
		min, max := 10, 30
		for i := 0; i < 10000; i++ {
			num := dashrand.Num(min, max)
			recv := num >= min && num < max
			if !assert.Equal(t, recv, true) {
				break
			}
		}
	}
	{
		min, max := int64(-10), int64(30)
		for i := 0; i < 10000; i++ {
			num := dashrand.Num(min, max)
			recv := num >= min && num < max
			if !assert.Equal(t, recv, true) {
				break
			}
		}
	}
	{
		min, max := int64(10), int64(30)
		for i := 0; i < 10000; i++ {
			num := dashrand.Num(min, max)
			recv := num >= min && num < max
			if !assert.Equal(t, recv, true) {
				break
			}
		}
	}
	{
		min, max := -12.345, 34.567
		for i := 0; i < 10000; i++ {
			num := dashrand.Num(min, max)
			recv := num >= min && num < max
			if !assert.Equal(t, recv, true) {
				break
			}
		}
	}
	{
		min, max := 12.345, 34.567
		for i := 0; i < 10000; i++ {
			num := dashrand.Num(min, max)
			recv := num >= min && num < max
			if !assert.Equal(t, recv, true) {
				break
			}
		}
	}
}

/*
TestIntSli is a ...
*/
func TestIntSli(t *testing.T) {
	lenArr := []uint{1, 3, 9, 13, 16, 32}
	for _, leng := range lenArr {
		{
			recv := dashrand.NumSlice(leng, -10, 30)
			assert.Equal(t, leng, uint(len(recv)))
		}
		{
			recv := dashrand.NumSlice(leng, 10, 30)
			assert.Equal(t, leng, uint(len(recv)))
		}
		{
			recv := dashrand.NumSlice(leng, -12.345, 34.567)
			assert.Equal(t, leng, uint(len(recv)))
		}
		{
			recv := dashrand.NumSlice(leng, 12.345, 34.567)
			assert.Equal(t, leng, uint(len(recv)))
		}
	}
}
