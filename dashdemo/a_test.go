package dashdemo_test

import (
	"log"
	"os"
	"testing"
)

/*
TestMain 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。
*/
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

/*
setup pre-test 安装
*/
func setup() {
	log.Println("Before all tests")
}

/*
teardown post-test 卸载
*/
func teardown() {
	log.Println("After all tests")
}
