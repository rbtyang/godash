package filedash_test

import (
	"github.com/rbtyang/godash/filedash"
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
