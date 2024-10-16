package dashrand_test

import (
	"log"
	"math/rand"
	"sync"
	"testing"

	"github.com/bwmarrin/snowflake"
)

var (
	sidMap  sync.Map
	nodeMap sync.Map
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestSnowOne is a ...
*/
func TestSnowOne(t *testing.T) {
	{
		node, err := snowflake.NewNode(rand.Int63n(1024))
		if err != nil {
			t.Error(err)
		}
		id := node.Generate().Int64() //like: 1533348896611241984
		t.Log(id)
	}
}

//串行基准测试

/*
BenchmarkSnowSeri go test -bench=SnowSeri$
*/
func BenchmarkSnowSeri(b *testing.B) {
	idMap := map[int64]bool{}
	node, err := snowflake.NewNode(1)
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	for n := 0; n < b.N; n++ {
		id := node.Generate().Int64() //生成id
		if _, ok := idMap[id]; ok {
			b.Errorf("cnt:%d, %d conflict", len(idMap), id)
		} else {
			idMap[id] = true
		}
	}

	b.Logf("cnt:%d", len(idMap))
}

//并行基准测试

/*
BenchmarkSnowPara go test -bench=SnowPara$
*/
func BenchmarkSnowPara(b *testing.B) {
	node, err := snowflake.NewNode(1) //同个节点，还是存在冲突
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			id := node.Generate().Int64() //生成id
			if _, ok := sidMap.Load(id); ok {
				b.Errorf("%d conflict", id)
			} else {
				sidMap.Store(id, true)
			}
		}
	})

	b.Logf("done")
}

//并行基准测试

/*
BenchmarkSnowRandNode go test -bench=SnowRandNode$
*/
func BenchmarkSnowRandNode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer() //b.ResetTimer()之前的处理 不会放到 执行时间里，也不会输出到报告中，所以可以在之前 做一些不计划 作为测试报告的操作

	b.RunParallel(func(pb *testing.PB) {
		var node *snowflake.Node
		nid := rand.Int63n(1024)
		if val, ok := nodeMap.Load(nid); ok {
			node = val.(*snowflake.Node)
		} else {
			var err error
			node, err = snowflake.NewNode(nid)
			if err != nil {
				b.Error(err)
			}
			nodeMap.Store(nid, node)
		}

		for pb.Next() {
			id := node.Generate().Int64() //生成id
			if _, ok := sidMap.Load(id); ok {
				b.Errorf("%d conflict", id)
			} else {
				sidMap.Store(id, true)
			}
		}
	})

	b.Logf("done")
}
