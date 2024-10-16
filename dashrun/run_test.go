package dashrun_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashrun"
	"github.com/rbtyang/godash/dashrun/internal/runtest"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

/*
@Editor robotyang at 2024

TestPanicTrace is a ...
*/
func TestPanicTrace(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			msg := fmt.Sprint(e)
			assert.Equal(t, "TestPanic[11:14 22:19 22:28]", msg)
			stacks := string(dashrun.PanicTrace(1))
			stacks = strings.ReplaceAll(stacks, "/", string(os.PathSeparator))

			i := strings.Index(stacks, `github.com\rbtyang\godash\dashrun_test.testPanicTrace2`)
			j := strings.Index(stacks, `github.com\rbtyang\godash\dashrun_test.testPanicTrace1`)
			k := strings.Index(stacks, `github.com\rbtyang\godash\dashrun_test.TestPanicTrace`)

			if i == -1 || j == -1 || k == -1 {
				assert.Fail(t, "not found stack path", "strings.Index=", i, j, k)
			}
			if i > j || j > k {
				assert.Fail(t, "stack sequence error", "strings.Index=", i, j, k)
			}
		}
	}()
	var data []string
	testPanicTrace1(t, data)
}

/*
@Editor robotyang at 2024

TestLastCallerPlace is a ...
*/
func TestLastCallerPlace(t *testing.T) {
	var data []string
	runtest.TestLastCallerPlace(t, data)
}

func testPanicTrace1(t *testing.T, data []string) []string {
	data = append(data, "11:14")
	data = testPanicTrace2(t, data)
	data = append(data, "11:16")
	return data
}

func testPanicTrace2(t *testing.T, data []string) []string {
	data = append(data, "22:19")
	data = append(data, "22:28")
	log.Panic("TestPanic", data)
	return data
}
