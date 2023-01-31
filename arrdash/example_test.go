package arrdash_test

import (
	"fmt"
	"github.com/rbtyang/godash/arrdash"
)

func ExampleInclude_slice_general() {
	var data1 = []int{8, 5, 5}
	recv := arrdash.Include(data1, 8) //here here here
	fmt.Println(recv)                 //bool
}

func ExampleInclude_slice_struct() {
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
	var data = []*Coder{coderZS, coderLS}
	{
		recv := arrdash.Include(data, coderLS) //here here here
		fmt.Println(recv)                      //bool
	}
}
