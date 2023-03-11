package dashfile_test

import (
	"github.com/rbtyang/godash/dasharr"
	"github.com/rbtyang/godash/dashfile"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
TestGetFilePathList is a ...

@Editor robotyang at 2023
*/
func TestGetFilePathList(t *testing.T) {
	{
		pathList, err := dashfile.GetFileList("../")
		if err != nil {
			t.Error(err)
		}
		t.Log(pathList)
	}
	{
		pathList, err := dashfile.GetFileListExcept("../", []string{".git", "vendor"})
		if err != nil {
			t.Error(err)
		}
		t.Log(pathList)
	}
}

/*
TestGetFileListFilter is a ...

@Editor robotyang at 2023
*/
func TestGetFileListFilter(t *testing.T) {
	filelist, err := dashfile.GetFileListFilter("../", []string{".go"}, []string{".git", "vendor", "_test.go"})
	if err != nil {
		t.Fatal(err)
	}
	{
		want := true
		recv := dasharr.Contain(filelist, "..\\dasharr\\arr.go")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dasharr.Contain(filelist, []string{
			"..\\dasharr\\rbtyang.go",
		})
		assert.Equal(t, want, recv)
	}
}
