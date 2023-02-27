package demodash_test

import (
	"github.com/rbtyang/godash/demodash"
	"github.com/rbtyang/godash/logdash"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
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
TestErrGroup is a ...

@Editor robotyang at 2023
*/
func TestErrGroup(t *testing.T) {
	var erg errgroup.Group

	erg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", demodash.SomeFunc("ZhangSan"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "LiSi Ni Hao", demodash.SomeFunc("LiSi"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "WangWu Ni Hao", demodash.SomeFunc("WangWu"))
		return nil
	})

	err := erg.Wait()
	if err != nil {
		logdash.Error(err)
	}
}
