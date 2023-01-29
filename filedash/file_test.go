package filedash_test

import (
	"github.com/rbtyang/godash/arrdash"
	"github.com/rbtyang/godash/filedash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFilePathList(t *testing.T) {
	{
		pathList, err := filedash.GetFileList("../")
		if err != nil {
			t.Error(err)
		}
		t.Log(pathList)
	}
	{
		pathList, err := filedash.GetFileListExcept("../", []string{".git", "vendor"})
		if err != nil {
			t.Error(err)
		}
		t.Log(pathList)
	}
}

func TestGetFileListFilter(t *testing.T) {
	filelist, err := filedash.GetFileListFilter("../", []string{".go"}, []string{".git", "vendor", "_test.go"})
	if err != nil {
		t.Fatal(err)
	}
	{
		want := true
		recv := arrdash.Contain(filelist, "..\\arrdash\\arr.go")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := arrdash.Contain(filelist, []string{
			"..\\arrdash\\rbtyang.go",
		})
		assert.Equal(t, want, recv)
	}
}
