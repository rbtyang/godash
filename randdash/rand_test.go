package randdash_test

import (
	"github.com/rbtyang/godash/randdash"
	"log"
	"testing"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before this tests")
}

/*
TestRandStr is a ...

@Editor robotyang at 2023
*/
func TestRandStr(t *testing.T) {
	{
		recv := randdash.Str(randdash.ModeNum, 32)
		t.Log(recv)
	}
	{
		recv := randdash.Str(randdash.Modeaz, 32)
		t.Log(recv)
	}
	{
		recv := randdash.Str(randdash.ModeNum+randdash.Modeaz, 32)
		t.Log(recv)
	}
	{
		recv := randdash.Str(randdash.ModeNum+randdash.ModeAZ, 32)
		t.Log(recv)
	}
	{
		recv := randdash.Str(randdash.ModeNumAlpha, 32)
		t.Log(recv)
	}
	{
		recv := randdash.Str(randdash.ModeNumAlphaSp, 32)
		t.Log(recv)
	}
}

/*
TestRandNum is a ...

@Editor robotyang at 2023
*/
func TestRandNum(t *testing.T) {
	{
		recv := randdash.Num(10, 30)
		t.Log(recv)
	}
}

/*
TestRandCode is a ...

@Editor robotyang at 2023
*/
func TestRandCode(t *testing.T) {
	{
		recv := randdash.NumCode(6)
		t.Log(recv)
	}
}

/*
TestRandIntSli is a ...

@Editor robotyang at 2023
*/
func TestRandIntSli(t *testing.T) {
	{
		recv := randdash.IntSli(100, 10, 30)
		t.Log(recv)
	}
}
