package errdash_test

import (
	"github.com/rbtyang/godash/errdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestTemp(t *testing.T) {
	msg := errdash.TransCodeMsg(99999999)
	assert.Equal(t, msg, "")
}

func TestTemp2(t *testing.T) {
	msg := errdash.TransCodeMsg(errdash.ErrParam)
	assert.Equal(t, msg, "参数错误")
}

func TestTemp3(t *testing.T) {
	errdash.RegisterDict(map[uint32]string{333333: "ha哈哈哈"})
	msg := errdash.TransCodeMsg(333333)
	assert.Equal(t, msg, "ha哈哈哈")
}
