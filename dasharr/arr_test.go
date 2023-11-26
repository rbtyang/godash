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
			output := dasharr.Include(input, 8)
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Include(input, 9)
			assert.Equal(t, want, output)
		}
	}
	//slice float32
	{
		var input = []float32{8.1, 5.23, 5.456}
		{
			want := true
			output := dasharr.Include(input, float32(5.456))
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Include(input, float32(9))
			assert.Equal(t, want, output)
		}
	}
	//slice string
	{
		var input = []string{"aaa", "bbb", "ccc"}
		{
			want := true
			output := dasharr.Include(input, "aaa")
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Include(input, "ddd")
			assert.Equal(t, want, output)
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
			output := dasharr.Include(input, coderLS)
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, 8)
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, 9)
			assert.Equal(t, want, output)
		}
	}
	//array int L2
	{
		var input = [...][3]int{{8, 5, 5}, {9, 5, 5}}
		{
			want := true
			output := dasharr.Contain(input, [3]int{9, 5, 5})
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, [3]int{9, 9, 6})
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, output)
		}
	}

	//slice int
	{
		var input = []int{8, 5, 5}
		{
			want := true
			output := dasharr.Contain(input, 8)
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, 9)
			assert.Equal(t, want, output)
		}
	}
	//slice float32
	{
		var input = []float32{8.1, 5.23, 5.456}
		{
			want := true
			output := dasharr.Contain(input, float32(5.456))
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, float32(9))
			assert.Equal(t, want, output)
		}
	}
	//slice string
	{
		input := []string{"rbtyang", "robotyang", "大绵羊"}
		{
			want := true
			output := dasharr.Contain(input, "rbtyang")
			assert.Equal(t, want, output)
		}
		{
			want := true
			output := dasharr.Contain(input, "robotyang")
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, "jackma")
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, "asdf")
			//output := arrdash.Contain(input, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, output)
		}
		{
			want := true
			output := dasharr.Contain(input, "ghjk")
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, "asdxxx")
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, coderZS)
			//output := arrdash.Contain(input, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, output)
		}
		{
			want := true
			output := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, coderQL)
			assert.Equal(t, want, output)
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
			output := dasharr.Contain(input, coderZS)
			//output := arrdash.Contain(input, map[string]string{"123":"asdfasfasd"})
			assert.Equal(t, want, output)
		}
		{
			want := true
			output := dasharr.Contain(input, coderLS)
			assert.Equal(t, want, output)
		}
		{
			want := false //-
			output := dasharr.Contain(input, coderQL)
			assert.Equal(t, want, output)
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
			output := dasharr.JoinAny(input, ",")
			assert.Equal(t, want, output)
		}
		{
			input := []interface{}{"rbt yang", "num", 9527}
			want := "rbt yang,num,9527"
			output := dasharr.JoinAny(input, ",")
			assert.Equal(t, want, output)
		}
	}
}

/*
Test_Chunk is a ...
*/
func Test_Chunk(t *testing.T) {
	{
		slice := []string{"a", "b", "c", "d"}
		{
			want := [][]string{{"a"}, {"b"}, {"c"}, {"d"}}
			output := dasharr.Chunk(slice, 1)
			assert.Equal(t, want, output)
		}
		{
			want := [][]string{{"a", "b", "c"}, {"d"}}
			output := dasharr.Chunk(slice, 3)
			assert.Equal(t, want, output)
		}
	}
	{
		slice := []any{"a", "b", 1, 2, "d"}
		{
			want := [][]any{{"a", "b", 1}, {2, "d"}}
			output := dasharr.Chunk(slice, 3)
			assert.Equal(t, want, output)
		}
	}
	{
		slice := []string{"a", "b", "c", "d", "d", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
		{
			want := [][]string{{"a", "b", "c"}, {"d", "d", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n"}}
			output := dasharr.Chunk(slice, 3)
			assert.Equal(t, want, output)
		}
		{
			want := [][]string{{"a", "b", "c", "d", "d", "f"}, {"g", "h", "i", "j", "k", "l"}, {"m", "n"}}
			output := dasharr.Chunk(slice, 6)
			assert.Equal(t, want, output)
		}
	}
}

/*
Benchmark_Chunk_RandIntArr is a ...
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
		output = dasharr.FilterBy(input, func(item int) bool {
			return item > 0
		})
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})

	dashtime.CostPrint("FilterByWg", time.Nanosecond, func() {
		output = dasharr.FilterByWg(input, func(item int) bool {
			return item > 0
		})
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})

	dashtime.CostPrint("FilterByChan", time.Nanosecond, func() {
		output = dasharr.FilterByChan(input, func(item int) bool {
			return item > 0
		}, 5)
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})
}

/*
Test_Filter_NilSlice 测试当传入的切片为空时，是否返回空切片。
*/
func Test_Filter_NilSlice(t *testing.T) {
	var input []int
	var output []int

	dashtime.CostPrint("FilterBy", time.Nanosecond, func() {
		output = dasharr.FilterBy(input, func(item int) bool {
			return item > 0
		})
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})

	dashtime.CostPrint("FilterByWg", time.Nanosecond, func() {
		output = dasharr.FilterByWg(input, func(item int) bool {
			return item > 0
		})
		assert.Equal(t, 0, len(output))
		assert.Equal(t, input, output)
	})

	dashtime.CostPrint("FilterByChan", time.Nanosecond, func() {
		output = dasharr.FilterByChan(input, func(item int) bool {
			return item > 0
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
		output1 := dasharr.FilterBy(input, func(item int) bool {
			return item > 0
		})
		output2 := dasharr.FilterBy(input, func(item int) bool {
			return item <= 0
		})
		output0 := append(output1, output2...)

		want := dashsort.Quick(input)
		output := dashsort.Quick(output0)
		assert.Equal(t, want, output)
	})

	dashtime.CostPrint("FilterByWg", time.Nanosecond, func() {
		output1 := dasharr.FilterByWg(input, func(item int) bool {
			return item > 0
		})
		output2 := dasharr.FilterByWg(input, func(item int) bool {
			return item <= 0
		})
		output0 := append(output1, output2...)

		want := dashsort.Quick(input)
		output := dashsort.Quick(output0)
		assert.Equal(t, want, output)
	})

	dashtime.CostPrint("FilterByChan", time.Nanosecond, func() {
		output1 := dasharr.FilterByChan(input, func(item int) bool {
			return item > 0
		}, 10000)
		output2 := dasharr.FilterByChan(input, func(item int) bool {
			return item <= 0
		}, 10000)
		output0 := append(output1, output2...)

		want := dashsort.Quick(input)
		output := dashsort.Quick(output0)
		assert.Equal(t, want, output)
	})
}

/*
Test_Filter_NilSlice 测试当传入的切片为空时，是否返回空切片。
*/
func Test_FilterNull(t *testing.T) {
	{
		input := []int{1, 2, 0, 0, 3, 0, 4, 5}
		dashtime.CostPrint("FilterByWg-FilterNull", time.Nanosecond, func() {
			want := []int{1, 2, 3, 4, 5}
			output := dasharr.FilterNull(input)
			assert.Equal(t, want, dashsort.Quick(output))
		})
	}
	{
		input := []string{"1a", "2b", "0", "0", "3c", "0", "4d", "5f"}
		dashtime.CostPrint("FilterByWg-FilterNull", time.Nanosecond, func() {
			want := []string{"1a", "2b", "3c", "4d", "5e"}
			output := dasharr.FilterNull(input) //lastedit
			assert.Equal(t, want, output)
		})
	}
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
