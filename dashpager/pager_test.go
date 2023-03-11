package dashpager_test

import (
	"github.com/rbtyang/godash/dashpager"
	"github.com/stretchr/testify/assert"
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
TestPager is a ...

@Editor robotyang at 2023
*/
func TestPager(t *testing.T) {
	{
		pager := dashpager.NewPagerAll(1, 10, true)
		assert.Equal(t, 1, pager.Index())
		assert.Equal(t, 10, pager.Size())
		assert.Equal(t, true, pager.NeedAll())
		assert.NotEqual(t, 11, pager.Size())
		assert.NotEqual(t, false, pager.NeedAll())
	}
	{
		pager := dashpager.NewPager(1, 10)
		assert.Equal(t, 1, pager.Index())
		assert.Equal(t, 10, pager.Size())
		assert.Equal(t, false, pager.NeedAll())
		assert.NotEqual(t, 11, pager.Size())
		assert.NotEqual(t, true, pager.NeedAll())
	}
}
