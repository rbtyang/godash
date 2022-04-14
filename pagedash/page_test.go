package pagedash_test

import (
	"github.com/rbtyang/godash/pagedash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestPager(t *testing.T) {
	{
		pager := pagedash.NewPagerAll(1, 10, true)
		assert.Equal(t, 1, pager.Index())
		assert.Equal(t, 10, pager.Size())
		assert.Equal(t, true, pager.NeedAll())
		assert.NotEqual(t, 11, pager.Size())
		assert.NotEqual(t, false, pager.NeedAll())
	}
	{
		pager := pagedash.NewPager(1, 10)
		assert.Equal(t, 1, pager.Index())
		assert.Equal(t, 10, pager.Size())
		assert.Equal(t, false, pager.NeedAll())
		assert.NotEqual(t, 11, pager.Size())
		assert.NotEqual(t, true, pager.NeedAll())
	}
}
