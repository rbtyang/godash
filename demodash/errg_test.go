package demodash_test

import (
	"github.com/rbtyang/godash/demodash"
	"github.com/rbtyang/godash/logdash"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
)

func TestErrg(t *testing.T) {
	var eg errgroup.Group

	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", demodash.HelloWorld("ZhangSan"))
		return nil
	})
	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao" , demodash.HelloWorld("LiSi"))
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
