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

func TestErrGroup(t *testing.T) {
	var eg errgroup.Group

	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", demodash.SomeFunc("ZhangSan"))
		return nil
	})
	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", demodash.SomeFunc("LiSi"))
		return nil
	})
	eg.Go(func() error {
		assert.Equal(t, "ZhangSan Bu Hao", demodash.SomeFunc("WangWu"))
		return nil
	})

	err := eg.Wait()
	if err != nil {
		logdash.Error(err)
	}
}
