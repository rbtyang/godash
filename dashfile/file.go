package dashfile

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// IsExist 判断 文件或文件夹 是否存在；
// 如果返回的 错误 为nil,说明文件或文件夹存在；
// 如果返回的 错误类型 使用os.IsNotExist() 判断为true, 说明 文件或文件夹 不存在；

/*
IsExist  如果返回的 错误 为其它类型, 则不确定 是否在存在；

@Editor robotyang at 2023
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
IsDir  IsDir 判断 所给路径 是否为文件夹

@Editor robotyang at 2023
*/
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/*
IsFile  IsFile 判断 所给路径 是否为文件

@Editor robotyang at 2023
*/
func IsFile(path string) bool {
	return !IsDir(path)
}

/*
IsExistFile 判断 文件 是否存在；

@Editor robotyang at 2023
*/
func IsExistFile(filePath string) bool {
	return IsExist(filePath) && IsFile(filePath)
}

/*
IsExistDir 判断 文件夹 是否存在；

@Editor robotyang at 2023
*/
func IsExistDir(dirPath string) bool {
	return IsExist(dirPath) && IsDir(dirPath)
}

/*
Rebuild 重建文件（会自动创建目录，文件存在则清空，不存在则创建）

@Editor robotyang at 2023
*/
func Rebuild(filePath string) (*os.File, error) {
	dirPath, _ := filepath.Split(filePath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, err
	}

	return os.Create(filePath)
}

/*
RebuildOrOpen  RebuildOrOpen 重建或打开文件（会自动创建目录，文件存在则打开，不存在则创建）

@Editor robotyang at 2023
*/
func RebuildOrOpen(filePath string) (*os.File, error) {
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

//ReadByFilePath 读取文件内容

/*
ReadByFilePath @reference https://haicoder.net/golang/golang-read.html

@Editor robotyang at 2023
*/
func ReadByFilePath(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ReadByFile(file)
}

//ReadByFile 读取文件内容

/*
ReadByFile @reference https://haicoder.net/golang/golang-read.html

@Editor robotyang at 2023
*/
func ReadByFile(file *os.File) ([]byte, error) {
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	return chunk, nil
}

/*
CopyFile  CopyFile 从一个文件 拷贝到 另一个文件

@Editor robotyang at 2023
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
CompareFileBySum  CompareFileBySum 通过计算文件签名，判断两个文件是否一致

@Editor robotyang at 2023
*/
func CompareFileBySum(filePath1, filePath2 string) bool {
	return calculateFileSum(filePath1) == calculateFileSum(filePath2)
}

// calculateFileSum 计算文件校验签名

/*
calculateFileSum  @reference https://blog.csdn.net/neweastsun/article/details/123666235

@Editor robotyang at 2023
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
GetLastDir 获取 最后一个 目录名

@Editor robotyang at 2023
*/
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

/*
GetLastDirWithCheck  GetLastDirWithCheck 先检查 目录是否存在，存在则 获取 最后一个 目录名

@Editor robotyang at 2023
*/
func GetLastDirWithCheck(path string) (string, error) {
	if !IsExist(path) || !IsDir(path) {
		return "", errors.New(path + "非有效目录路径")
	}
	return GetLastDir(path), nil
}

/*
GetFileList is a ...

@Editor robotyang at 2023
*/
func GetFileList(path string) ([]string, error) {
	return getFileList(path, nil, nil)
}

/*
GetFileListFilter is a ...

@Editor robotyang at 2023
*/
func GetFileListFilter(path string, accept, except []string) ([]string, error) {
	return getFileList(path, accept, except)
}

/*
GetFileListAccept is a ...

@Editor robotyang at 2023
*/
func GetFileListAccept(path string, accept []string) ([]string, error) {
	return getFileList(path, accept, nil)
}

/*
GetFileListExcept is a ...

@Editor robotyang at 2023
*/
func GetFileListExcept(path string, except []string) ([]string, error) {
	return getFileList(path, nil, except)
}

//getFileList 方法名
//@param path 根目录

/*
getFileList @param except 排除的 目录、文件 的关键词

@Editor robotyang at 2023
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
		for _, acpt := range accept { //包含的
			if !strings.Contains(path, acpt) {
				return nil
			}
		}
		for _, ecpt := range except { //排除的
			if strings.Contains(path, ecpt) {
				return nil
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
