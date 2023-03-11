package dashconvert_test

import (
	"github.com/rbtyang/godash/dashconvert"
	"github.com/stretchr/testify/assert"
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
TestStrToByte is a ...

@Editor robotyang at 2023
*/
func TestStrToByte(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := dashconvert.StrToByte("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := "hello world 123 哈哈"
		recv := dashconvert.ByteToStr([]byte("hello world 123 哈哈"))
		assert.Equal(t, want, recv)
	}
}

/*
TestStrToByteByUnsafe is a ...

@Editor robotyang at 2023
*/
func TestStrToByteByUnsafe(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := dashconvert.StrToByteByUnsafe("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := "hello world 123 哈哈"
		recv := dashconvert.ByteToStrByUnsafe([]byte("hello world 123 哈哈"))
		assert.Equal(t, want, recv)
	}
}

/*
TestStrToByteByReflect is a ...

@Editor robotyang at 2023
*/
func TestStrToByteByReflect(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := dashconvert.StrToByteByReflect("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
}
