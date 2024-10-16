package dashconv_test

import (
	"log"
	"testing"

	"github.com/rbtyang/godash/dashconv"
	"github.com/stretchr/testify/assert"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestStrToByte is a ...
*/
func TestStrToByte(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := dashconv.StrToByte("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := "hello world 123 哈哈"
		recv := dashconv.ByteToStr([]byte("hello world 123 哈哈"))
		assert.Equal(t, want, recv)
	}
}

/*
TestStrToByteByUnsafe is a ...
*/
func TestStrToByteByUnsafe(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := dashconv.StrToByteByUnsafe("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := "hello world 123 哈哈"
		recv := dashconv.ByteToStrByUnsafe([]byte("hello world 123 哈哈"))
		assert.Equal(t, want, recv)
	}
}

/*
TestStrToByteByReflect is a ...
*/
func TestStrToByteByReflect(t *testing.T) {
	{
		//want := []byte("hello world 123 哈哈")
		//recv := dashconv.StrToByteByReflect("hello world 123 哈哈")
		//assert.Equal(t, want, recv)
	}
}
