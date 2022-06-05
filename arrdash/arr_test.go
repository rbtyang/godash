package errdash_test

import (
	errdash "github.com/rbtyang/godash/arrdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestContains(t *testing.T) {
	// array
	{
		var data [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
		{
			want := true
			recv := errdash.Contains(data, [3]int{1, 2, 3})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := errdash.Contains(data, [3]int{1, 2, 5})
			assert.NotEqual(t, want, recv)
		}
	}

	// slice
	{
		data := []string{"123", "asd", "11xx"}
		{
			want := true
			recv := errdash.Contains(data, "123")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := errdash.Contains(data, "asd")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := errdash.Contains(data, "asdxxx")
			assert.NotEqual(t, want, recv)
		}
	}

	// map
	{
		data := map[string]string{
			"123": "qwer",
			"asd": "asdf",
			"asd3": "asdf",
			"11asd": "zxcv",
		}
		{
			want := true
			recv := errdash.Contains(data, "qwer")
			//recv := errdash.Contains(data, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := errdash.Contains(data, "asdf")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := errdash.Contains(data, "asdxxx")
			assert.NotEqual(t, want, recv)
		}
	}

}

func TestArrayToString(t *testing.T) {
	{
		data := []interface{}{"asdfasdf", "1123323", "sdf23134"}
		{
			want := "asdfasdf,1123323,sdf23134"
			recv := errdash.ArrayToString(data)
			assert.Equal(t, want, recv)
		}
	}
}
