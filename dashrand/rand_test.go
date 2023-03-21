package dashrand_test

import (
	"github.com/rbtyang/godash/dashrand"
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
		recv := dashrand.Num(10, 30)
		t.Log(recv)
	}
}

/*
@Editor robotyang at 2023

TestRandCode is a ...
*/
func TestRandCode(t *testing.T) {
	{
		recv := dashrand.NumCode(6)
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
