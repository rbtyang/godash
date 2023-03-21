package dashfile

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
@Editor robotyang at 2023

IsExist 判断 文件或文件夹 是否存在；
*/
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

/*
@Editor robotyang at 2023

IsDir 判断 所给路径 是否存在 且为文件夹
*/
func IsDir(path string) bool {
	if !IsExist(path) {
		return false
	}
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/*
@Editor robotyang at 2023

IsFile 判断 所给路径 是否存在 且为文件
*/
func IsFile(path string) bool {
	if !IsExist(path) {
		return false
	}
	return !IsDir(path)
}

/*
@Editor robotyang at 2023

OpenFile 打开文件（是 os.OpenFile 别名）
*/
func OpenFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, os.ModePerm)
}

/*
@Editor robotyang at 2023

CreateOrOpen 重建或打开文件（会自动创建目录，文件存在则打开，不存在则创建）
*/
func CreateOrOpen(filePath string) (*os.File, error) {
	dirPath, _ := filepath.Split(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}

	if IsExist(filePath) {
		return os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, os.ModePerm)
	} else {
		return os.Create(filePath)
	}
}

/*
@Editor robotyang at 2023

CreateOrReset 重建文件（会自动创建目录，文件存在则清空，不存在则创建）
*/
func CreateOrReset(filePath string) (*os.File, error) {
	dirPath, _ := filepath.Split(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

/*
@Editor robotyang at 2023

ReadByFilePath 读取文件内容

@Editor robotyang at 2023

@Reference https://haicoder.net/golang/golang-read.html
*/
func ReadByFilePath(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadByFile(file)
}

/*
@Editor robotyang at 2023

ReadByFile 读取文件内容
*/
func ReadByFile(file *os.File) ([]byte, error) {
	return ioutil.ReadAll(file)
}

/*
@Editor robotyang at 2023

RemoveAll 删除文件或目录

在删除文件时，RemoveAll()和Remove()方法没有太大的区别；

在删除目录时，Remove()只能删除空目录，而RemoveAll()不受任何限制。
*/
func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

/*
@Editor robotyang at 2023

CopyFile 从src文件 读取内容 拷贝到 dst文件

@Param srcFilePath 来源文件路径

@Param dstFilePath 目标文件路径：文件存在则覆盖，不存在则创建
*/
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

/*
@Editor robotyang at 2023

CopyFile 从src文件 读取内容 追加到 dst文件

@Param srcFilePath 来源文件路径

@Param dstFilePath 目标文件路径：文件存在则覆盖，不存在则创建
*/
func AppendFile(srcFilePath, dstFilePath string) (writeen int, err error) {
	srcCont, err := ReadByFilePath(srcFilePath)
	if err != nil {
		return 0, err
	}
	dstFile, err := OpenFile(dstFilePath)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()
	return dstFile.WriteString(string(srcCont))
}

/*
@Editor robotyang at 2023

CompareFileBySum 通过计算文件签名，判断两个文件是否一致
*/
func CompareFileBySum(filePath1, filePath2 string) bool {
	return calculateFileSum(filePath1) == calculateFileSum(filePath2)
}

// calculateFileSum 计算文件校验签名

/*
@Editor robotyang at 2023

calculateFileSum

@Editor robotyang at 2023

@Reference https://blog.csdn.net/neweastsun/article/details/123666235
*/
func calculateFileSum(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	//h := sha256.New()
	//h := sha1.New()
	//h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	// 格式化为16进制字符串
	return fmt.Sprintf("%x", h.Sum(nil))
}

/*
@Editor robotyang at 2023

GetLastDir 获取 最后一个 目录名
*/
func GetLastDir(path string) string {
	sep := string(os.PathSeparator)
	path2, _ := filepath.Abs(path)
	if strings.HasSuffix(path, "\\") ||
		strings.HasSuffix(path, "/") ||
		strings.HasSuffix(path, ".") {
		path2 += string(os.PathSeparator)
	}
	dir, _ := filepath.Split(path2)
	dir = strings.TrimRight(dir, "\\/")      //去掉最后的 /
	dir = strings.ReplaceAll(dir, "\\", sep) //转换所有的 /
	dirArr := strings.Split(dir, sep)        //根据 / 进行分割
	lastDir := dirArr[len(dirArr)-1]         //获取 最后一个 目录名
	return lastDir
}

/*
@Editor robotyang at 2023

GetLastDirWithCheck 先检查路径资源是否存在，存在则获取 最后一个目录名
*/
func GetLastDirWithCheck(path string) (string, error) {
	if !IsExist(path) {
		return "", errors.New("非有效资源路径")
	}
	return GetLastDir(path), nil
}

/*
@Editor robotyang at 2023

GetFileList is a ...
*/
func GetFileList(path string) ([]string, error) {
	return getFileList(path, nil, nil)
}

/*
@Editor robotyang at 2023

GetFileListFilter 递归获取目录下的所有子文件（排除清单except > 保留清单accept）
*/
func GetFileListFilter(path string, accept, except []string) ([]string, error) {
	return getFileList(path, accept, except)
}

/*
@Editor robotyang at 2023

GetFileListAccept is a ...
*/
func GetFileListAccept(path string, accept []string) ([]string, error) {
	return getFileList(path, accept, nil)
}

/*
@Editor robotyang at 2023

GetFileListExcept is a ...
*/
func GetFileListExcept(path string, except []string) ([]string, error) {
	return getFileList(path, nil, except)
}

/*
@Editor robotyang at 2023

getFileList

@Param except 排除的 目录、文件 的关键词
*/
func getFileList(path string, accept, except []string) ([]string, error) {
	var pathList []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() { //目录的
			return nil
		}

		if len(accept) > 0 {
			var acptIt bool
			for _, acpt := range accept { //包含的
				if strings.Contains(path, acpt) {
					acptIt = true
					break
				}
			}
			if !acptIt {
				return nil
			}
		}

		if len(except) > 0 {
			for _, ecpt := range except { //排除的
				if strings.Contains(path, ecpt) {
					return nil
				}
			}
		}

		pathList = append(pathList, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return pathList, nil
}
