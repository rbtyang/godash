package dashdemo_test

import (
	"github.com/rbtyang/godash/dashdemo"
	"github.com/rbtyang/godash/dashlog"
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
		assert.Equal(t, "ZhangSan Ni Hao", dashdemo.SomeFunc("ZhangSan"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "LiSi Ni Hao", dashdemo.SomeFunc("LiSi"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "WangWu Ni Hao", dashdemo.SomeFunc("WangWu"))
		return nil
	})

	err := erg.Wait()
	if err != nil {
		dashlog.Error(err)
	}
}
