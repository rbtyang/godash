package dashrand_test

import (
	"github.com/google/uuid"
	"log"
	"sync"
	"testing"
)

var (
	uidMap sync.Map
)

/*
@Editor robotyang at 2023

init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
@Editor robotyang at 2023

TestUuidOne is a ...
*/
func TestUuidOne(t *testing.T) {
	{
		uid := uuid.New().String()
		t.Log(uid)
	}
}

//串行基准测试

/*
@Editor robotyang at 2023

BenchmarkUuidSeri go test -bench=UuidSeri$
*/
func BenchmarkUuidSeri(b *testing.B) {
	idMap := map[string]bool{}

	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		id := uuid.New().String()
		if _, ok := idMap[id]; ok {
			b.Errorf("cnt:%d, %s conflict", len(idMap), id)
		} else {
			idMap[id] = true
		}
	}

	b.Logf("cnt:%d", len(idMap))
}

//并行基准测试

/*
@Editor robotyang at 2023

BenchmarkUuidRandNode go test -bench=UuidRandNode$
*/
func BenchmarkUuidRandNode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			id := uuid.New().String()
			if _, ok := uidMap.Load(id); ok {
				b.Errorf("%s conflict", id)
			} else {
				uidMap.Store(id, true)
			}
		}
	})

	b.Logf("done")
}
