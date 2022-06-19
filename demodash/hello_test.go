package demodash_test

import (
	"github.com/rbtyang/godash/demodash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestHelloWorld(t *testing.T) {
	{
		want := "ZhangSan Ni Hao"
		recv := demodash.HelloWorld("ZhangSan")
		assert.Equal(t, want, recv)
	}
	{
		want := "LiSi Ni Hao"
		recv := demodash.HelloWorld("ZhangSan")
		assert.Equal(t, want, recv)
	}
}

//go test -bench=.  //. 是全部
//go test -bench=HelloWorld$  //正则匹配
//可选参数：
//  -cpu 1,2,4 指定运行的cpu 格式
//  -count n 指定运行的次数
//  -benchtime 每一条测试执行的时间 （默认是1s）
//  -bench 指定执行bench的方法， . 是全部
//  -benchmem 显示内存分配情况
//  其他参数 可以通过 go help testflag 查看
func BenchmarkHelloWorld(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		demodash.HelloWorld("ZhangSan" + string(n))
	}
}
