package dashrand_test

import (
	"github.com/rbtyang/godash/dashrand"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/*
@Editor robotyang at 2023

init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
@Editor robotyang at 2023

TestRandStr is a ...
*/
func TestRandStr(t *testing.T) {
	{
		recv := dashrand.Str(dashrand.ModeNum, 32)
		t.Log(recv)
	}
	{
		recv := dashrand.Str(dashrand.Modeaz, 32)
		t.Log(recv)
	}
	{
		recv := dashrand.Str(dashrand.ModeNum+dashrand.Modeaz, 32)
		t.Log(recv)
	}
	{
		recv := dashrand.Str(dashrand.ModeNum+dashrand.ModeAZ, 32)
		t.Log(recv)
	}
	{
		recv := dashrand.Str(dashrand.ModeNumAlpha, 32)
		t.Log(recv)
	}
	{
		recv := dashrand.Str(dashrand.ModeNumAlphaSp, 32)
		t.Log(recv)
	}
}

/*
@Editor robotyang at 2023

TestRandNum is a ...
*/
func TestRandNum(t *testing.T) {
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
@Editor robotyang at 2023

TestRandCode is a ...
*/
func TestRandCode(t *testing.T) {
	{
		recv := dashrand.NumLen(6)
		t.Log(recv)
	}
}

/*
@Editor robotyang at 2023

TestRandIntSli is a ...
*/
func TestRandIntSli(t *testing.T) {
	{
		recv := dashrand.IntSli(100, 10, 30)
		t.Log(recv)
	}
}
