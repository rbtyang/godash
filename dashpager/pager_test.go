package dashpager_test

import (
	"log"
	"testing"

	"github.com/rbtyang/godash/dashpager"
	"github.com/stretchr/testify/assert"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestPager is a ...
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
