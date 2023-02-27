package errdash_test

import (
	"github.com/rbtyang/godash/errdash"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
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

/*
TestParseCode is a ...

@Editor robotyang at 2023
*/
func TestParseCode(t *testing.T) {
	{
		want := errdash.CodeInternal
		recv := errdash.ParseCode(nil)
		assert.Equal(t, want, recv)
	}

	{
		want := errdash.CodeAlreadyExists
		recv := errdash.ParseCode(codes.AlreadyExists)
		assert.Equal(t, want, recv)
	}

	{
		want := errdash.CodeInternal
		recv := errdash.ParseCode("哈哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := errdash.CodePermissionDenied
		recv := errdash.ParseCode(7)
		assert.Equal(t, want, recv)
	}
	{
		want := errdash.CodePermissionDenied
		recv := errdash.ParseCode("7")
		assert.Equal(t, want, recv)
	}
	{
		want := uint32(99999)
		recv := errdash.ParseCode(99999)
		assert.Equal(t, want, recv)
	}
	{
		want := uint32(99999)
		recv := errdash.ParseCode("99999")
		assert.Equal(t, want, recv)
	}
}

/*
TestGetCodeMsg is a ...

@Editor robotyang at 2023
*/
func TestGetCodeMsg(t *testing.T) {
	{
		want := ""
		msg := errdash.GetCodeMsg(99999)
		assert.Equal(t, msg, want)
	}
	{
		want := "无效参数"
		recv := errdash.GetCodeMsg(errdash.CodeInvalidArgument)
		assert.Equal(t, want, recv)
	}
	{
		want := "ha哈哈哈"
		errdash.RegisterCode(map[uint32]string{88888: "ha哈哈哈"})
		recv := errdash.GetCodeMsg(88888)
		assert.Equal(t, want, recv)
	}

}
