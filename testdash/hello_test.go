package testdash_test

import (
	"github.com/rbtyang/godash/testdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Init this tests")
}

func TestHello(t *testing.T) {
	{
		want := "rbtyang你好，欢迎使用godash开发工具集"
		recv := testdash.Hello("rbtyang")
		assert.Equal(t, want, recv)
	}
}
