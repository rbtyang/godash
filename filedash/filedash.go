package filedash

import (
	"io"
	"os"
	"fmt"
)

/*
 判断文件或文件夹是否存在；
 如果返回的错误为nil,说明文件或文件夹存在；
 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在；
 如果返回的错误为其它类型,则不确定是否在存在；
*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*
 从一个文件拷贝到另一个文件；
*/
func CopyFile(dstName, srcName string) (writeen int64, err error) {
	src, err := os.Open(dstName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(srcName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}