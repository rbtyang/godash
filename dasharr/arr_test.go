package dasharr_test

//go:generate go test -v -run="none"  -bench=. -benchmem ./...

import (
	"github.com/rbtyang/godash/dasharr"
	"github.com/rbtyang/godash/dashrand"
	"github.com/rbtyang/godash/dashsort"
	"github.com/rbtyang/godash/dashtime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/*
Test_SliceHas is a ...
*/
func Test_SliceHas(t *testing.T) {
	//slice int
	{
		var input = []int{8, 5, 5}
		{
			want := true
			recv := dasharr.Include(input, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Include(input, 9)
			assert.Equal(t, want, recv)
		}
	}
	//slice float32
	{
		var input = []float32{8.1, 5.23, 5.456}
		{
			want := true
			recv := dasharr.Include(input, float32(5.456))
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Include(input, float32(9))
			assert.Equal(t, want, recv)
		}
	}
	//slice string
	{
		var input = []string{"aaa", "bbb", "ccc"}
		{
			want := true
			recv := dasharr.Include(input, "aaa")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Include(input, "ddd")
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
		var input = []*Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Include(input, coderLS)
			assert.Equal(t, want, recv)
		}
	}
}

/*
Test_Contains is a ...
*/
func Test_Contains(t *testing.T) {
	//array int L1
	{
		var input = [...]int{8, 5, 5}
		{
			want := true
			recv := dasharr.Contain(input, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, 9)
			assert.Equal(t, want, recv)
		}
	}
	//array int L2
	{
		var input = [...][3]int{{8, 5, 5}, {9, 5, 5}}
		{
			want := true
			recv := dasharr.Contain(input, [3]int{9, 5, 5})
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, [3]int{9, 9, 6})
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
		var input = [...]Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Contain(input, coderLS)
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
		var input = [...]*Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, recv)
		}
	}

	//slice int
	{
		var input = []int{8, 5, 5}
		{
			want := true
			recv := dasharr.Contain(input, 8)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, 9)
			assert.Equal(t, want, recv)
		}
	}
	//slice float32
	{
		var input = []float32{8.1, 5.23, 5.456}
		{
			want := true
			recv := dasharr.Contain(input, float32(5.456))
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, float32(9))
			assert.Equal(t, want, recv)
		}
	}
	//slice string
	{
		input := []string{"rbtyang", "robotyang", "大绵羊"}
		{
			want := true
			recv := dasharr.Contain(input, "rbtyang")
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(input, "robotyang")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, "jackma")
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
		var input = []Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Contain(input, coderLS)
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
		var input = []*Coder{coderZS, coderLS, coderWW}
		{
			want := true
			recv := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, recv)
		}
	}

	//map string
	{
		input := map[string]string{
			"123": "asdf",
			"qwe": "ghjk",
			"rty": "zxcv",
			"uio": "bnmla",
		}
		{
			want := true
			recv := dasharr.Contain(input, "asdf")
			//recv := arrdash.Contain(input, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(input, "ghjk")
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, "asdxxx")
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
		input := map[string]Coder{
			"123": coderZS,
			"qwe": coderLS,
			"rty": coderWW,
		}
		{
			want := true
			recv := dasharr.Contain(input, coderZS)
			//recv := arrdash.Contain(input, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, coderQL)
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
		input := map[string]*Coder{
			"123": coderZS,
			"qwe": coderLS,
			"rty": coderWW,
		}
		{
			want := true
			recv := dasharr.Contain(input, coderZS)
			//recv := arrdash.Contain(input, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, recv)
		}
		{
			want := false //-
			recv := dasharr.Contain(input, coderQL)
			assert.Equal(t, want, recv)
		}
	}
}

/*
Test_ArrayToString is a ...
*/
func Test_ArrayToString(t *testing.T) {
	{
		{
			input := []interface{}{"rbtyang", "num", 9527}
			want := "rbtyang,num,9527"
			recv := dasharr.JoinAny(input, ",")
			assert.Equal(t, want, recv)
		}
		{
			input := []interface{}{"rbt yang", "num", 9527}
			want := "rbt yang,num,9527"
			recv := dasharr.JoinAny(input, ",")
			assert.Equal(t, want, recv)
		}
	}
}

/*
Test_Chunk is a ...
*/
func Test_Chunk(t *testing.T) {
	{
		array := []string{"a", "b", "c", "d"}
		{
			want := [][]string{{"a", "b", "c"}, {"d"}}
			recv := dasharr.Chunk(array, 0)
			assert.Equal(t, want, recv)
		}
		{
			want := [][]string{{"a", "b", "c"}, {"d"}}
			recv := dasharr.Chunk(array, 3)
			assert.Equal(t, want, recv)
		}
	}
}

/*
Test_Chunk_RandIntArr is a ...
*/
func Test_Chunk_RandIntArr(t *testing.T) {
	//input := dashrand.NumSlice(100000, -100, 100)
	{
		array := []string{"a", "b", "c", "d"}
		{
			want := [][]string{{"a"}, {"b"}, {"c"}, {"d"}}
			recv := dasharr.Chunk(array, 1)
			assert.Equal(t, want, recv)
		}
		{
			want := [][]string{{"a", "b", "c"}, {"d"}}
			recv := dasharr.Chunk(array, 3)
			assert.Equal(t, want, recv)
		}
	}
	{
		array := []string{"a", "b", "c", "d", "d", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
		{
			want := [][]string{{"a", "b", "c"}, {"d", "d", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n"}}
			recv := dasharr.Chunk(array, 3)
			assert.Equal(t, want, recv)
		}
		{
			want := [][]string{{"a", "b", "c", "d", "d", "f"}, {"g", "h", "i", "j", "k", "l"}, {"m", "n"}}
			recv := dasharr.Chunk(array, 6)
			assert.Equal(t, want, recv)
		}
	}
}

/*
Test_Chunk_RandIntArr is a ...
*/
func Benchmark_Chunk_RandIntArr(b *testing.B) {
	input := dashrand.NumSlice(100000, -100, 100)
	for n := 0; n < b.N; n++ {
		dasharr.Chunk(input, uint(dashrand.Num(1, 100000)))
	}
}

/*
Test_Filter_EmptySlice 测试当传入的切片为空时，是否返回空切片。
*/
func Test_Filter_EmptySlice(t *testing.T) {
	input := []int{}
	var output []int

	dashtime.CostPrint("FilterBy", time.Nanosecond, func() {
		output = dasharr.FilterBy(input, func(int) bool {
			return true
		})
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})

	dashtime.CostPrint("FilterByWg", time.Nanosecond, func() {
		output = dasharr.FilterByWg(input, func(int) bool {
			return true
		})
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})

	dashtime.CostPrint("FilterByChan", time.Nanosecond, func() {
		output = dasharr.FilterByChan(input, func(int) bool {
			return true
		}, 5)
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})
}

/*
Test_Filter_RandIntArr 测试过滤随机整数切片
*/
func Test_Filter_RandIntArr(t *testing.T) {
	input := dashrand.NumSlice(100000, -100, 100)

	dashtime.CostPrint("FilterBy", time.Nanosecond, func() {
		recv1 := dasharr.FilterBy(input, func(item int) bool {
			return item > 0
		})
		recv2 := dasharr.FilterBy(input, func(item int) bool {
			return item <= 0
		})
		recv0 := append(recv1, recv2...)

		want := dashsort.Quick(input)
		recv := dashsort.Quick(recv0)
		assert.Equal(t, want, recv)
	})

	dashtime.CostPrint("FilterByWg", time.Nanosecond, func() {
		recv1 := dasharr.FilterByWg(input, func(item int) bool {
			return item > 0
		})
		recv2 := dasharr.FilterByWg(input, func(item int) bool {
			return item <= 0
		})
		recv0 := append(recv1, recv2...)

		want := dashsort.Quick(input)
		recv := dashsort.Quick(recv0)
		assert.Equal(t, want, recv)
	})

	dashtime.CostPrint("FilterByChan", time.Nanosecond, func() {
		recv1 := dasharr.FilterByChan(input, func(item int) bool {
			return item > 0
		}, 10000)
		recv2 := dasharr.FilterByChan(input, func(item int) bool {
			return item <= 0
		}, 10000)
		recv0 := append(recv1, recv2...)

		want := dashsort.Quick(input)
		recv := dashsort.Quick(recv0)
		assert.Equal(t, want, recv)
	})
}

/*
Benchmark_Filter_RandIntArr_FilterBy 串行的基准测试

@Reference https://blog.csdn.net/bestzy6/article/details/125515985

@Reference https://dandelioncloud.cn/article/details/1568817400504348673
*/
func Benchmark_Filter_RandIntArr_FilterBy(b *testing.B) {
	input := dashrand.NumSlice(100000, -100, 100)
	b.ResetTimer() //在开始基准测试前，可能有一些很耗时的准备工作，这个时候就可以通过b.ResetTimer() 重新开始计时
	for n := 0; n < b.N; n++ {
		dasharr.FilterBy(input, func(item int) bool {
			return item > 0
		})
		dasharr.FilterBy(input, func(item int) bool {
			return item <= 0
		})
	}
}

/*
Benchmark_Filter_RandIntArr_FilterByWg 串行的基准测试
*/
func Benchmark_Filter_RandIntArr_FilterByWg(b *testing.B) {
	input := dashrand.NumSlice(100000, -100, 100)
	b.ResetTimer() //在开始基准测试前，可能有一些很耗时的准备工作，这个时候就可以通过b.ResetTimer() 重新开始计时
	for n := 0; n < b.N; n++ {
		dasharr.FilterByWg(input, func(item int) bool {
			return item > 0
		})
		dasharr.FilterByWg(input, func(item int) bool {
			return item <= 0
		})
	}
}

/*
Benchmark_Filter_RandIntArr_FilterByChan 串行的基准测试
*/
func Benchmark_Filter_RandIntArr_FilterByChan(b *testing.B) {
	input := dashrand.NumSlice(100000, -100, 100)
	b.ResetTimer() //在开始基准测试前，可能有一些很耗时的准备工作，这个时候就可以通过b.ResetTimer() 重新开始计时
	for n := 0; n < b.N; n++ {
		dasharr.FilterByChan(input, func(item int) bool {
			return item > 0
		}, 1000)
		dasharr.FilterByChan(input, func(item int) bool {
			return item <= 0
		}, 1000)
	}
}
