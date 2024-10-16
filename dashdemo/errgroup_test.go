package dashdemo_test

import (
	"log"
	"testing"

	"github.com/rbtyang/godash/dashdemo"
	"github.com/rbtyang/godash/dashlog"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
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

ErrGroup is a ...
*/
func TestErrGroup(t *testing.T) {
	var erg errgroup.Group

	erg.Go(func() error {
		assert.Equal(t, "ZhangSan Ni Hao", dashdemo.ErrGroup("ZhangSan"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "LiSi Ni Hao", dashdemo.ErrGroup("LiSi"))
		return nil
	})
	erg.Go(func() error {
		assert.Equal(t, "WangWu Ni Hao", dashdemo.ErrGroup("WangWu"))
		return nil
	})

	err := erg.Wait()
	if err != nil {
		dashlog.Error(err)
	}
}
