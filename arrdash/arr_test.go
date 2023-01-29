package arrdash_test

import (
	"github.com/rbtyang/godash/arrdash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceHas(t *testing.T) {
	//slice int
	{
		var data = []int{8, 5, 5}
		{
			want := true
			recv := arrdash.SliceHas(data, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.SliceHas(data, 9)
			assert.Equal(t, want, recv)
		}
	}
	//slice float32
	{
		var data = []float32{8.1, 5.23, 5.456}
		{
			want := true
			recv := arrdash.SliceHas(data, float32(5.456))
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.SliceHas(data, float32(9))
			assert.Equal(t, want, recv)
		}
	}
	//slice string
	{
		var data = []string{"aaa", "bbb", "ccc"}
		{
			want := true
			recv := arrdash.SliceHas(data, "aaa")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.SliceHas(data, "ddd")
			assert.Equal(t, want, recv)
		}
	}
	//slice *struct
	{
		type Coder struct {
			Name  string
			Hobby []string
		}
		coderZS := &Coder{
			Name:  "ZhangSan",
			Hobby: []string{"唱", "跳"},
		}
		coderLS := &Coder{
			Name:  "LiSi",
			Hobby: []string{"rap", "篮球"},
		}
		coderWW := &Coder{
			Name:  "WangWu",
			Hobby: []string{"吃饭", "睡觉", "打豆豆"},
		}
		var data = []*Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := arrdash.SliceHas(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}
}

func TestContains(t *testing.T) {
	//array L1
	{
		var data = [...]int{8, 5, 5}
		{
			want := true
			recv := arrdash.Contain(data, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.Contain(data, 9)
			assert.Equal(t, want, recv)
		}
	}
	//array L2
	{
		var data = [...][3]int{{8, 5, 5}, {9, 5, 5}}
		{
			want := true
			recv := arrdash.Contain(data, [3]int{9, 5, 5})
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.Contain(data, [3]int{9, 9, 6})
			assert.Equal(t, want, recv)
		}
	}

	//slice int
	{
		var data = []int{8, 5, 5}
		{
			want := true
			recv := arrdash.Contain(data, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.Contain(data, 9)
			assert.Equal(t, want, recv)
		}
	}
	//slice float32
	{
		var data = []float32{8.1, 5.23, 5.456}
		{
			want := true
			recv := arrdash.Contain(data, float32(5.456))
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.Contain(data, float32(9))
			assert.Equal(t, want, recv)
		}
	}
	//slice string
	{
		data := []string{"rbtyang", "robotyang", "大绵羊"}
		{
			want := true
			recv := arrdash.Contain(data, "rbtyang")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := arrdash.Contain(data, "robotyang")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.Contain(data, "jackma")
			assert.Equal(t, want, recv)
		}
	}

	//slice struct
	{
		type Coder struct {
			Name  string
			Hobby []string
		}
		coderZS := Coder{
			Name:  "ZhangSan",
			Hobby: []string{"唱", "跳"},
		}
		coderLS := Coder{
			Name:  "LiSi",
			Hobby: []string{"rap", "篮球"},
		}
		coderWW := Coder{
			Name:  "WangWu",
			Hobby: []string{"吃饭", "睡觉", "打豆豆"},
		}
		var data = []Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := arrdash.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}
	//slice *struct
	{
		type Coder struct {
			Name  string
			Hobby []string
		}
		coderZS := &Coder{
			Name:  "ZhangSan",
			Hobby: []string{"唱", "跳"},
		}
		coderLS := &Coder{
			Name:  "LiSi",
			Hobby: []string{"rap", "篮球"},
		}
		coderWW := &Coder{
			Name:  "WangWu",
			Hobby: []string{"吃饭", "睡觉", "打豆豆"},
		}
		var data = []*Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := arrdash.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}

	//map
	{
		data := map[string]string{
			"123":   "qwer",
			"asd":   "asdf",
			"asd3":  "asdf",
			"11asd": "zxcv",
		}
		{
			want := true
			recv := arrdash.Contain(data, "qwer")
			//recv := arrdash.Contain(data, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := arrdash.Contain(data, "asdf")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := arrdash.Contain(data, "asdxxx")
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
