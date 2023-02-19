package sortdash_test

import (
	"fmt"
	"github.com/rbtyang/godash/randdash"
	"github.com/rbtyang/godash/sortdash"
	"github.com/rbtyang/godash/timedash"
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

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

func TestBubble(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		sortdash.Bubble(data)
		assert.Equal(t, want, data)
	}
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		sortdash.Bubble2(data)
		assert.Equal(t, want, data)
	}
}

func TestInsertion(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := timedash.Cost(func() {
			sortdash.Insertion(data)
			assert.Equal(t, want, data)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := timedash.Cost(func() {
			sortdash.Insertion2(data)
			assert.Equal(t, want, data)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
}

func TestSelection(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		sortdash.Selection(data)
		assert.Equal(t, want, data)
	}
}

func TestQuick(t *testing.T) {
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := timedash.Cost(func() {
			recv := sortdash.Quick(data)
			assert.Equal(t, want, recv)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
	{
		var data = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}
		want := []int{1, 21, 32, 36, 39, 43, 54, 62, 66, 76, 78}
		cost := timedash.Cost(func() {
			recv := sortdash.QuickParallel(data)
			assert.Equal(t, want, recv)
		})
		fmt.Printf("cost: %d\n", cost.Nanoseconds())
	}
}

//go test -bench=QuickSimple$
func BenchmarkQuickSimple(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		randSli := randdash.IntSli(100, 0, 90)
		sortdash.Quick(randSli)
	}
}

//go test -bench=QuickParallel$
func BenchmarkQuickParallel(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		randSli := randdash.IntSli(100, 0, 90)
		sortdash.QuickParallel(randSli)
	}
}
