package runtest

import (
	"runtime"
	"strings"
	"testing"

	"github.com/rbtyang/godash/dashrun"
	"github.com/stretchr/testify/assert"
)

func TestLastCallerPlace(t *testing.T, data []string) []string {
	data = append(data, "11:14")
	data = testLastCallerPlace1(t, data)
	data = append(data, "11:16")
	return data
}

func testLastCallerPlace1(t *testing.T, data []string) []string {
	data = append(data, "11:14")
	data = testLastCallerPlace2(t, data)
	data = append(data, "11:16")
	return data
}

func testLastCallerPlace2(t *testing.T, data []string) []string {
	data = append(data, "22:19")
	data = append(data, "22:28")

	lst := dashrun.LastCallerFuncName(true)
	assert.Equal(t, "runtest.testLastCallerPlace1", lst)
	lstFull := dashrun.LastCallerFuncName(false)
	assert.Equal(t, "github.com/rbtyang/godash/dashrun/internal/runtest.testLastCallerPlace1", lstFull)

	_, curFileName, _, _ := runtime.Caller(0)
	lstPath := dashrun.LastCallerPlace(curFileName)
	bol := strings.Contains(lstPath, "godash/dashrun/run_test.go")
	assert.Equal(t, true, bol)

	return data
}
