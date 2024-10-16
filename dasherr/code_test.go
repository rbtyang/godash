package dasherr_test

import (
	"log"
	"testing"

	"github.com/rbtyang/godash/dasherr"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestParseCode is a ...
*/
func TestRegisterCode(t *testing.T) {
	dasherr.RegisterCode(map[uint32]string{
		19001: "自定义Code1",
		19002: "自定义Code2",
	})
	{
		want := "自定义Code1"
		recv := dasherr.GetCodeMsg(19001)
		assert.Equal(t, want, recv)
	}
	{
		want := "自定义Code2"
		recv := dasherr.GetCodeMsg(19002)
		assert.Equal(t, want, recv)
	}
}

/*
TestParseCode is a ...
*/
func TestParseCode(t *testing.T) {
	{
		want := dasherr.CodeInternal
		recv := dasherr.ParseCode(nil)
		assert.Equal(t, want, recv)
	}
	{
		want := dasherr.CodeAlreadyExists
		recv := dasherr.ParseCode(codes.AlreadyExists)
		assert.Equal(t, want, recv)
	}
	{
		want := dasherr.CodeInternal
		recv := dasherr.ParseCode("哈哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := dasherr.CodePermissionDenied
		recv := dasherr.ParseCode(7)
		assert.Equal(t, want, recv)
	}
	{
		want := dasherr.CodePermissionDenied
		recv := dasherr.ParseCode("7")
		assert.Equal(t, want, recv)
	}
	{
		want := uint32(99999)
		recv := dasherr.ParseCode(99999)
		assert.Equal(t, want, recv)
	}
	{
		want := uint32(99999)
		recv := dasherr.ParseCode("99999")
		assert.Equal(t, want, recv)
	}
}

/*
TestGetCodeMsg is a ...
*/
func TestGetCodeMsg(t *testing.T) {
	{
		want := "未知错误"
		msg := dasherr.GetCodeMsg(99999)
		assert.Equal(t, msg, want)
	}
	{
		want := "无效参数"
		recv := dasherr.GetCodeMsg(dasherr.CodeInvalidArgument)
		assert.Equal(t, want, recv)
	}
	{
		want := "ha哈哈哈"
		dasherr.RegisterCode(map[uint32]string{88888: "ha哈哈哈"})
		recv := dasherr.GetCodeMsg(88888)
		assert.Equal(t, want, recv)
	}

}
