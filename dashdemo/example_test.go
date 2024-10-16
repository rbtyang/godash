package dashdemo_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashdemo"
	"github.com/rbtyang/godash/dashlog"
	"golang.org/x/sync/errgroup"
	"sort"
)

/*
go并发编程示例

Example of go concurrent programming
*/
func ExampleErrGroup_demo() {
	var erg errgroup.Group
	var results []string // 用于收集结果的切片

	erg.Go(func() error {
		results = append(results, dashdemo.ErrGroup("ZhangSan"))
		return nil
	})
	erg.Go(func() error {
		results = append(results, dashdemo.ErrGroup("LiSi"))
		return nil
	})
	erg.Go(func() error {
		results = append(results, dashdemo.ErrGroup("WangWu"))
		return nil
	})

	err := erg.Wait()
	if err != nil {
		dashlog.Error(err)
	}

	sort.Strings(results)
	fmt.Println(results)
	//Output: [LiSi Ni Hao WangWu Ni Hao ZhangSan Ni Hao]
}
