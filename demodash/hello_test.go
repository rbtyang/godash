package demodash_test

import (
	"github.com/rbtyang/godash/demodash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestHelloWorld(t *testing.T) {
	{
		want := "ZhangSan Ni Hao"
		recv := demodash.HelloWorld("ZhangSan")
		assert.Equal(t, want, recv)
	}
	{
		want := "LiSi Ni Hao"
		recv := demodash.HelloWorld("ZhangSan")
		assert.Equal(t, want, recv)
	}
}