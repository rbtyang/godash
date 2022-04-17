package demodash_test

import (
	"github.com/rbtyang/godash/demodash"
	"github.com/rbtyang/godash/logdash"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
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

func TestHelloWorldByErrg(t *testing.T) {
	var eg errgroup.Group

	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", demodash.HelloWorld("ZhangSan"))
		return nil
	})
	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", demodash.HelloWorld("LiSi"))
		return nil
	})
	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Bu Hao", demodash.HelloWorld("WangWu"))
		return nil
	})

	err := eg.Wait()
	if err != nil {
		logdash.Error(err)
	}
}
