package dashrandom_test

import (
	"github.com/rbtyang/godash/dashrandom"
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
		recv := dashrandom.Str(dashrandom.ModeNum, 32)
		t.Log(recv)
	}
	{
		recv := dashrandom.Str(dashrandom.Modeaz, 32)
		t.Log(recv)
	}
	{
		recv := dashrandom.Str(dashrandom.ModeNum+dashrandom.Modeaz, 32)
		t.Log(recv)
	}
	{
		recv := dashrandom.Str(dashrandom.ModeNum+dashrandom.ModeAZ, 32)
		t.Log(recv)
	}
	{
		recv := dashrandom.Str(dashrandom.ModeNumAlpha, 32)
		t.Log(recv)
	}
	{
		recv := dashrandom.Str(dashrandom.ModeNumAlphaSp, 32)
		t.Log(recv)
	}
}

/*
TestRandNum is a ...

@Editor robotyang at 2023
*/
func TestRandNum(t *testing.T) {
	{
		recv := dashrandom.Num(10, 30)
		t.Log(recv)
	}
}

/*
TestRandCode is a ...

@Editor robotyang at 2023
*/
func TestRandCode(t *testing.T) {
	{
		recv := dashrandom.NumCode(6)
		t.Log(recv)
	}
}

/*
TestRandIntSli is a ...

@Editor robotyang at 2023
*/
func TestRandIntSli(t *testing.T) {
	{
		recv := dashrandom.IntSli(100, 10, 30)
		t.Log(recv)
	}
}
