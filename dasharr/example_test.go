package dasharr_test

import (
	"fmt"
	"github.com/rbtyang/godash/dasharr"
)

/*
通用切片入参示例

Example of common slice entry parameter
*/
func ExampleInclude_slice_general() {
	var data1 = []int{8, 5, 5}
	recv := dasharr.Include(data1, 8) //here
	fmt.Println(recv)
	// Output: true
}

/*
结构体切片入参示例

Example of structure slice entry parameter
*/
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
		recv := dasharr.Include(data, coderLS) //here
		fmt.Println(recv)
	}
	//Output: true
}
