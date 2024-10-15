package dashfile_test

import (
	"github.com/rbtyang/godash/dasharr"
	"github.com/rbtyang/godash/dashfile"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

/*
TestIsExist is a ...
*/
func TestIsExist(t *testing.T) {
	pwd, _ := os.Getwd()
	t.Log(pwd)
	{
		want := true
		recv := dashfile.IsExist("./file.go")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsExist("./file111.go")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashfile.IsExist("../dashfile")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsExist("../dashfile1111")
		assert.Equal(t, want, recv)
	}
}

/*
TestIsDir is a ...
*/
func TestIsDir(t *testing.T) {
	{
		want := false
		recv := dashfile.IsDir("./file.go")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsDir("./file111.go")
		assert.Equal(t, want, recv)
	}
	{
		want := true
		recv := dashfile.IsDir("../dashfile")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsDir("../dashfile1111")
		assert.Equal(t, want, recv)
	}
}

/*
TestIsFile is a ...
*/
func TestIsFile(t *testing.T) {
	{
		want := true
		recv := dashfile.IsFile("./file.go")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsFile("./file111.go")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsFile("../dashfile")
		assert.Equal(t, want, recv)
	}
	{
		want := false
		recv := dashfile.IsFile("../dashfile1111")
		assert.Equal(t, want, recv)
	}
}

func TestCreateOrOpen(t *testing.T) {
	path := "../temp/CreateOrOpen.txt"
	err := dashfile.RemoveAll(path)
	if err != nil {
		t.Error(err)
		return
	}

	file, err := dashfile.CreateOrOpen(path)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = file.WriteString("你好，世界！\n你好，中国~\n")
	if err != nil {
		t.Error(err)
		return
	}
	file.Close()

	file2, err := dashfile.CreateOrOpen(path)
	if err != nil {
		t.Error(err)
		return
	}
	oldCont, err := dashfile.ReadByFile(file2)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "你好，世界！\n你好，中国~\n", string(oldCont))

	_, err = file2.WriteString("你好，Golang！\n")
	if err != nil {
		t.Error(err)
		return
	}
	file2.Close()

	oldCont2, err := dashfile.ReadByFilePath(path)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "你好，世界！\n你好，中国~\n你好，Golang！\n", string(oldCont2))
}

func TestCreateOrReset(t *testing.T) {
	path := "../temp/CreateOrReset.txt"
	err := dashfile.RemoveAll(path)
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 2; i++ {
		file, err := dashfile.CreateOrReset(path)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = file.WriteString("你好，世界！\n你好，中国~\n")
		if err != nil {
			t.Error(err)
			return
		}
		file.Close()

		oldCont2, err := dashfile.ReadByFilePath(path)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, "你好，世界！\n你好，中国~\n", string(oldCont2))
	}
}

func TestCopyFile(t *testing.T) {
	srcPath := "../temp/CopyFileSrc.txt"
	dstPath := "../temp/CopyFileDst.txt"

	{
		path := srcPath
		file, err := dashfile.CreateOrReset(path)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = file.WriteString("你好，世界！\n你好，中国~\n")
		if err != nil {
			t.Error(err)
			return
		}
		file.Close()
	}

	{
		_, err := dashfile.CopyFile(srcPath, dstPath)
		if err != nil {
			t.Error(err)
			return
		}
	}

	{
		path := dstPath
		cont, err := dashfile.ReadByFilePath(path)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, "你好，世界！\n你好，中国~\n", string(cont))
	}

}

func TestAppendFile(t *testing.T) {
	srcPath := "../temp/AppendFileSrc.txt"
	dstPath := "../temp/AppendFileDst.txt"

	{
		path := srcPath
		file, err := dashfile.CreateOrReset(path)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = file.WriteString("你好，世界！\n你好，中国~\n")
		if err != nil {
			t.Error(err)
			return
		}
		file.Close()
	}
	{
		path := dstPath
		file, err := dashfile.CreateOrReset(path)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = file.WriteString("哈喽，世界！\n哈喽，中国~\n")
		if err != nil {
			t.Error(err)
			return
		}
		file.Close()
	}

	{
		_, err := dashfile.AppendFile(srcPath, dstPath)
		if err != nil {
			return
		}
		cont, err := dashfile.ReadByFilePath(dstPath)
		if err != nil {
			t.Error(err)
			return
		}
		assert.Equal(t, "哈喽，世界！\n哈喽，中国~\n你好，世界！\n你好，中国~\n", string(cont))
	}
}

/*
TestGetLastDir is a ...
*/
func TestGetLastDir(t *testing.T) {
	{
		want := "dashfile"
		recv := dashfile.GetLastDir("./file.go") //存在的 合法格式文件
		assert.Equal(t, want, recv)
	}
	{
		want := "godash"
		recv := dashfile.GetLastDir("../fileaaa.go") //不存在的 合法格式文件
		assert.Equal(t, want, recv)
	}
	{
		want := "godash"
		recv := dashfile.GetLastDir("../dashaaa") //不存在的 未知格式文件
		assert.Equal(t, want, recv)
	}
	{
		want := "dasharr"
		recv := dashfile.GetLastDir("../dasharr/") //存在的 目录
		assert.Equal(t, want, recv)
	}
	{
		want := "dasharr"
		recv := dashfile.GetLastDir("../dasharr\\") //存在的 目录
		assert.Equal(t, want, recv)
	}
	{
		want := "dashaaa"
		recv := dashfile.GetLastDir("../dashaaa/") //不存在的 目录
		assert.Equal(t, want, recv)
	}
	{
		want := "dashaaa"
		recv := dashfile.GetLastDir("../dashaaa\\") //不存在的 目录
		assert.Equal(t, want, recv)
	}
	{
		want := "dashfile"
		recv := dashfile.GetLastDir(".") //当前目录
		assert.Equal(t, want, recv)
	}
	{
		want := "dashfile"
		recv := dashfile.GetLastDir("./") //当前目录
		assert.Equal(t, want, recv)
	}
}

/*
TestGetLastDir is a ...
*/
func TestGetLastDirWithCheck(t *testing.T) {
	{
		lastDir, err := dashfile.GetLastDirWithCheck("../dashaaa/fileaaa.go")
		assert.Equal(t, "非有效资源路径", err.Error())
		assert.Equal(t, "", lastDir)
	}
	{
		lastDir, err := dashfile.GetLastDirWithCheck("../dasharr/arr.go")
		assert.Equal(t, nil, err)
		assert.Equal(t, "dasharr", lastDir)
	}
}

/*
TestGetFilePathList is a ...
*/
func TestGetFileList(t *testing.T) {
	{
		pathList, err := dashfile.GetFileList("../")
		if err != nil {
			t.Error(err)
			return
		}
		assert.Greater(t, len(pathList), 70)
	}
	{
		pathList, err := dashfile.GetFileListAccept("../dasharr/", []string{"arr.go", "arr_test.go"})
		if err != nil {
			t.Error(err)
			return
		}
		want := []string{"..\\dasharr\\arr.go", "..\\dasharr\\arr_test.go"}
		assert.Equal(t, want, pathList)
	}
	{
		pathList, err := dashfile.GetFileListExcept("../", []string{".git", "vendor"})
		if err != nil {
			t.Error(err)
			return
		}
		assert.Greater(t, len(pathList), 60)
	}
	{
		pathList, err := dashfile.GetFileListFilter("../dasharr/",
			[]string{"arr.go", "arr_test.go"},         //保留清单
			[]string{".git", "vendor", "arr_test.go"}, //排除清单
		)
		if err != nil {
			t.Error(err)
			return
		}
		want := []string{"..\\dasharr\\arr.go"}
		assert.Equal(t, want, pathList)
	}
}

/*
TestGetFileListFilter is a ...
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
