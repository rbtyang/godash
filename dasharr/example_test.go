package dasharr_test

import (
	"fmt"
	"github.com/rbtyang/godash/dasharr"
)

/*
ExampleInclude_slice_general is a ...

@Editor robotyang at 2023
*/
func ExampleInclude_slice_general() {
	var data1 = []int{8, 5, 5}
	recv := dasharr.Include(data1, 8) //here here here
	fmt.Println(recv)                 //bool
}

/*
ExampleInclude_slice_struct is a ...

@Editor robotyang at 2023
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
		recv := dasharr.Include(data, coderLS) //here here here
		fmt.Println(recv)                      //bool
	}
}
