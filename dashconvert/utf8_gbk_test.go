package dashconvert_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashconvert"
	"log"
	"testing"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before this tests")
}

func TestUtf8ToGbk(t *testing.T) {
	s := "GBK 与 UTF-8 编码转换测试"

	gbk, err := dashconvert.Utf8ToGbk([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Utf8ToGbk: " + string(gbk))
	}

	utf8, err := dashconvert.GbkToUtf8(gbk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GbkToUtf8: " + string(utf8))
	}
}

/*
TestTransEncoding is a ...

@Editor robotyang at 2023
*/
func TestTransEncoding(t *testing.T) {
	s := "GBK 与 UTF-8 编码转换测试"
	gbk, _ := dashconvert.Utf8ToGbk([]byte(s))
	fmt.Println("Utf8ToGbk: " + string(gbk))

	utf8 := dashconvert.TransEncoding(string(gbk), "gbk", "utf-8")
	fmt.Println("TransToUtf8: " + string(utf8))
}
