package dasharr_test

import (
	"github.com/rbtyang/godash/dasharr"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
@Editor robotyang at 2023

TestSliceHas is a ...
*/
func TestSliceHas(t *testing.T) {
	//slice int
	{
		var data = []int{8, 5, 5}
		{
			want := true
			recv := dasharr.Include(data, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Include(data, 9)
			assert.Equal(t, want, recv)
		}
	}
	//slice float32
	{
		var data = []float32{8.1, 5.23, 5.456}
		{
			want := true
			recv := dasharr.Include(data, float32(5.456))
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Include(data, float32(9))
			assert.Equal(t, want, recv)
		}
	}
	//slice string
	{
		var data = []string{"aaa", "bbb", "ccc"}
		{
			want := true
			recv := dasharr.Include(data, "aaa")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Include(data, "ddd")
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
			recv := dasharr.Include(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}
}

/*
@Editor robotyang at 2023

TestContains is a ...
*/
func TestContains(t *testing.T) {
	//array int L1
	{
		var data = [...]int{8, 5, 5}
		{
			want := true
			recv := dasharr.Contain(data, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, 9)
			assert.Equal(t, want, recv)
		}
	}
	//array int L2
	{
		var data = [...][3]int{{8, 5, 5}, {9, 5, 5}}
		{
			want := true
			recv := dasharr.Contain(data, [3]int{9, 5, 5})
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, [3]int{9, 9, 6})
			assert.Equal(t, want, recv)
		}
	}
	//array struct
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
		var data = [...]Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}
	//array *struct
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
		var data = [...]*Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}

	//slice int
	{
		var data = []int{8, 5, 5}
		{
			want := true
			recv := dasharr.Contain(data, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, 9)
			assert.Equal(t, want, recv)
		}
	}
	//slice float32
	{
		var data = []float32{8.1, 5.23, 5.456}
		{
			want := true
			recv := dasharr.Contain(data, float32(5.456))
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, float32(9))
			assert.Equal(t, want, recv)
		}
	}
	//slice string
	{
		data := []string{"rbtyang", "robotyang", "大绵羊"}
		{
			want := true
			recv := dasharr.Contain(data, "rbtyang")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(data, "robotyang")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, "jackma")
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
			recv := dasharr.Contain(data, coderLS)
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
			recv := dasharr.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
	}

	//map string
	{
		data := map[string]string{
			"123": "asdf",
			"qwe": "ghjk",
			"rty": "zxcv",
			"uio": "bnmla",
		}
		{
			want := true
			recv := dasharr.Contain(data, "asdf")
			//recv := arrdash.Contain(data, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(data, "ghjk")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, "asdxxx")
			assert.Equal(t, want, recv)
		}
	}

	//map struct
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
		coderQL := Coder{
			Name:  "QianLiu",
			Hobby: []string{"吃饭", "睡觉", "搞王者"},
		}
		data := map[string]Coder{
			"123": coderZS,
			"qwe": coderLS,
			"rty": coderWW,
		}
		{
			want := true
			recv := dasharr.Contain(data, coderZS)
			//recv := arrdash.Contain(data, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, coderQL)
			assert.Equal(t, want, recv)
		}
	}
	//map *struct
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
		coderQL := &Coder{
			Name:  "QianLiu",
			Hobby: []string{"吃饭", "睡觉", "搞王者"},
		}
		data := map[string]*Coder{
			"123": coderZS,
			"qwe": coderLS,
			"rty": coderWW,
		}
		{
			want := true
			recv := dasharr.Contain(data, coderZS)
			//recv := arrdash.Contain(data, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(data, coderLS)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(data, coderQL)
			assert.Equal(t, want, recv)
		}
	}
}

/*
@Editor robotyang at 2023

TestArrayToString is a ...
*/
func TestArrayToString(t *testing.T) {
	{
		{
			data := []interface{}{"rbtyang", "num", 9527}
			want := "rbtyang,num,9527"
			recv := dasharr.JoinAny(data, ",")
			assert.Equal(t, want, recv)
		}
		{
			data := []interface{}{"rbt yang", "num", 9527}
			want := "rbt yang,num,9527"
			recv := dasharr.JoinAny(data, ",")
			assert.Equal(t, want, recv)
		}
	}
}
