package arrdash_test

import (
	"github.com/rbtyang/godash/arrdash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	// array
	{
		var data = [...][3]int{{8, 5, 5}, {9, 5, 5}}
		{
			want := true
			recv := arrdash.Contains(data, [3]int{9, 5, 5})
			assert.Equal(t, want, recv)
		}
		{
			want := false
			recv := arrdash.Contains(data, [3]int{9, 9, 6})
			assert.Equal(t, want, recv)
		}
	}

	// slice
	{
		data := []string{"rbtyang", "robotyang", "大绵羊"}
		{
			want := true
			recv := arrdash.Contains(data, "rbtyang")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := arrdash.Contains(data, "robotyang")
			assert.Equal(t, want, recv)
		}
		{
			want := false
			recv := arrdash.Contains(data, "jackma")
			assert.Equal(t, want, recv)
		}
	}

	// map
	{
		data := map[string]string{
			"123":   "qwer",
			"asd":   "asdf",
			"asd3":  "asdf",
			"11asd": "zxcv",
		}
		{
			want := true
			recv := arrdash.Contains(data, "qwer")
			//recv := arrdash.Contains(data, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := arrdash.Contains(data, "asdf")
			assert.Equal(t, want, recv)
		}
		{
			want := false
			recv := arrdash.Contains(data, "asdxxx")
			assert.Equal(t, want, recv)
		}
	}

}

func TestArrayToString(t *testing.T) {
	{
		{
			data := []interface{}{"rbtyang", "num", 9527}
			want := "rbtyang,num,9527"
			recv := arrdash.JoinAny(data, ",")
			assert.Equal(t, want, recv)
		}
		{
			data := []interface{}{"rbt yang", "num", 9527}
			want := "rbt yang,num,9527"
			recv := arrdash.JoinAny(data, ",")
			assert.Equal(t, want, recv)
		}
	}
}
