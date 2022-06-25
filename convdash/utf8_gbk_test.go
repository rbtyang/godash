package convdash_test

import (
	"fmt"
	"github.com/rbtyang/godash/convdash"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestUtf8ToGbk(t *testing.T) {
	s := "GBK 与 UTF-8 编码转换测试"

	gbk, err := convdash.Utf8ToGbk([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Utf8ToGbk: " + string(gbk))
	}

	utf8, err := convdash.GbkToUtf8(gbk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GbkToUtf8: " + string(utf8))
	}
}

func TestTransEncoding(t *testing.T) {
	s := "GBK 与 UTF-8 编码转换测试"
	gbk, _ := convdash.Utf8ToGbk([]byte(s))
	fmt.Println("Utf8ToGbk: " + string(gbk))

	utf8 := convdash.TransEncoding(string(gbk), "gbk", "utf-8")
	fmt.Println("TransToUtf8: " + string(utf8))
}
