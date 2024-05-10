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
@Editor robotyang at 2023

init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
@Editor robotyang at 2023

DebugErrGroup is a ...
*/
func TestErrGroup(t *testing.T) {
	var erg errgroup.Group

	erg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", dashdemo.DebugErrGroup("ZhangSan"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "LiSi Ni Hao", dashdemo.DebugErrGroup("LiSi"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "WangWu Ni Hao", dashdemo.DebugErrGroup("WangWu"))
		return nil
	})

	err := erg.Wait()
	if err != nil {
		dashlog.Error(err)
	}
}
