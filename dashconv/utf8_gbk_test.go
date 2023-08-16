package dashconv_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashconv"
	"log"
	"testing"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

func TestUtf8ToGbk(t *testing.T) {
	s := "GBK 与 UTF-8 编码转换测试"

	gbk, err := dashconv.Utf8ToGbk([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Utf8ToGbk: " + string(gbk))
	}

	utf8, err := dashconv.GbkToUtf8(gbk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GbkToUtf8: " + string(utf8))
	}
}

/*
TestTransEncoding is a ...
*/
func TestTransEncoding(t *testing.T) {
	s := "GBK 与 UTF-8 编码转换测试"
	gbk, _ := dashconv.Utf8ToGbk([]byte(s))
	fmt.Println("Utf8ToGbk: " + string(gbk))

	utf8 := dashconv.TransEncoding(string(gbk), "gbk", "utf-8")
	fmt.Println("TransToUtf8: " + string(utf8))
}
