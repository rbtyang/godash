package filedash

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 判断 所给路径 是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断 所给路径 是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 判断 文件或文件夹 是否存在；
// 如果返回的 错误 为nil,说明文件或文件夹存在；
// 如果返回的 错误类型 使用os.IsNotExist() 判断为true, 说明 文件或文件夹 不存在；
// 如果返回的 错误 为其它类型, 则不确定 是否在存在；
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// 判断 文件 是否存在；
func IsExistFile(filePath string) bool {
	return IsFile(filePath) && IsExist(filePath)
}

// 判断 文件夹 是否存在；
func IsExistDir(dirPath string) bool {
	return IsDir(dirPath) && IsExist(dirPath)
}

func Create(filePath string) (*os.File, error) {
	dirPath, _ := filepath.Split(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

func CreateIfNotExist(filePath string) (*os.File, error) {
	dirPath, _ := filepath.Split(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}

	if IsExist(filePath) {
		return os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	} else {
		return os.Create(filePath)
	}
}

// 从一个文件 拷贝到 另一个文件
func CopyFile(srcFilePath, dstFilePath string) (writeen int64, err error) {
	//打开 原文件
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()

	//打开或创建 目标文件
	dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()

	//复制 文件
	return io.Copy(dstFile, srcFile)
}

// 获取 最后一个 目录名
func GetLastDir(path string) string {
	sep := string(os.PathSeparator)
	path, _ = filepath.Abs(path)
	dir, _ := filepath.Split(path)
	dir = strings.TrimRight(dir, "\\/")      //去掉最后的 /
	dir = strings.ReplaceAll(dir, "\\", sep) //转换所有的 /
	dirArr := strings.Split(dir, sep)        //根据 / 进行分割
	lastDir := dirArr[len(dirArr)-1]         //获取 最后一个 目录名
	return lastDir
}

func GetLastDirWithCheck(path string) (string, error) {
	if !IsExist(path) || !IsDir(path) {
		return "", errors.New(path + "非有效目录路径")
	}
	return GetLastDir(path), nil
}
