package dashsort_test

import (
	"fmt"
	"log"
	"sort"
	"testing"

	"github.com/rbtyang/godash/dashrand"
	"github.com/rbtyang/godash/dashsort"
	"github.com/rbtyang/godash/dashtime"
	"github.com/stretchr/testify/assert"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestBuiltinSortSlice is a ...
*/
func TestBuiltinSortSlice(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		assert.Equal(t, want, data)
	}
	{
		type People struct {
			Name string
			Age  int
		}
		data := []People{
			{"Gopher", 7},
			{"Alice", 55},
			{"Vera", 24},
			{"Bob", 75},
		}
		want := []People{
			{"Alice", 55},
			{"Bob", 75},
			{"Gopher", 7},
			{"Vera", 24},
		}
		notWant := []People{
			{"Bob", 75},
			{"Alice", 55},
			{"Vera", 24},
			{"Gopher", 7},
		}
		{ //底层：快排排序
			sort.Slice(data, func(i, j int) bool { return data[i].Name < data[j].Name })
			assert.Equal(t, want, data)
			assert.NotEqual(t, notWant, data)
		}
		{ //底层：插入排序
			sort.SliceStable(data, func(i, j int) bool { return data[i].Name < data[j].Name })
			assert.Equal(t, want, data)
			assert.NotEqual(t, notWant, data)
		}
	}
}

/*
TestBubble is a ...
*/
func TestBubble(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		dashsort.Bubble(data)
		assert.Equal(t, want, data)
	}
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		dashsort.Bubble2(data)
		assert.Equal(t, want, data)
	}
}

/*
TestInsertion is a ...
*/
func TestInsertion(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := dashtime.Cost(func() {
			dashsort.Insertion(data)
			assert.Equal(t, want, data)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := dashtime.Cost(func() {
			dashsort.Insertion2(data)
			assert.Equal(t, want, data)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
}

/*
TestSelection is a ...
*/
func TestSelection(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		dashsort.Selection(data)
		assert.Equal(t, want, data)
	}
}

/*
TestQuick is a ...
*/
func TestQuick(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := dashtime.Cost(func() {
			recv := dashsort.Quick(data)
			assert.Equal(t, want, recv)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := dashtime.Cost(func() {
			recv := dashsort.QuickParallel(data)
			assert.Equal(t, want, recv)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
}

/*
BenchmarkQuickSimple go test -bench=QuickSimple$
*/
func BenchmarkQuickSimple(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		randSli := dashrand.NumSlice(100, 0, 90)
		dashsort.Quick(randSli)
	}
}

/*
BenchmarkQuickParallel go test -bench=QuickParallel$
*/
func BenchmarkQuickParallel(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		randSli := dashrand.NumSlice(100, 0, 90)
		dashsort.QuickParallel(randSli)
	}
}
