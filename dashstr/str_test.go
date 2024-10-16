package dashstr_test

import (
	"testing"

	"github.com/rbtyang/godash/dashstr"
	"github.com/stretchr/testify/assert"
)

func Test_TrimMobile(t *testing.T) {
	{
		input := "1\v2 3\f4　5\t6\r7\n8\r\n9"
		want := "123456789"
		recv := dashstr.TrimBlank(input)
		assert.Equal(t, want, recv)
	}
}
