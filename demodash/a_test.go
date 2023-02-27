package demodash_test

import (
	"log"
	"os"
	"testing"
)

/*
TestMain 如果测试文件中包含函数 TestMain，那么生成的测试将调用 TestMain(m)，而不是直接运行测试。

@Editor robotyang at 2023
*/
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

/*
setup pre-test 安装

@Editor robotyang at 2023
*/
func setup() {
	log.Println("Before all tests")
}

/*
teardown post-test 卸载

@Editor robotyang at 2023
*/
func teardown() {
	log.Println("After all tests")
}
