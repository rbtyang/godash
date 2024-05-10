package dashcrypt

import (
	"bufio"
	"os"
)

/*
FileEncryptByZyx @Editor robotyang at 2023

# FileEncryptByZyx 文件加密；

@Param srcFilePath：需加密的 文件路径；

@Param dstFilePath：加密后 文件路径；

@Param secret：密钥；

@Reference https://www.jianshu.com/p/0caab60fea9f
*/
func FileEncryptByZyx(srcFilePath, dstFilePath string, secret []byte) (err error) {
	f, err := os.Open(srcFilePath)
	if err != nil { //未找到文件
		return err
	}
	defer f.Close()

	fInfo, _ := f.Stat()
	fLen := fInfo.Size() //待处理文件大小

	maxLen := 1024 * 1024 * 100 //100mb  每 100mb 进行加密一次
	var forNum int64 = 0
	getLen := fLen

	if fLen > int64(maxLen) {
		getLen = int64(maxLen)
		forNum = fLen / int64(maxLen) //需要加密次数=forNum+1
	}

	//加密后存储的文件
	ff, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil { //文件写入错误
		return err
	}
	defer ff.Close()

	//循环加密，并且写入文件
	for i := 0; i < int(forNum+1); i++ {
		a := make([]byte, getLen)
		n, err := f.Read(a)
		if err != nil { //文件读取错误
			return err
		}

		getByte, err := AesBs64Encrypt(a[:n], secret)
		if err != nil { //加密错误
			return err
		}
		//换行处理，有点乱了，想到更好的再改
		getBytes := append([]byte(getByte), []byte("\n")...)

		//写入
		buf := bufio.NewWriter(ff)
		if _, err := buf.WriteString(string(getBytes[:])); err != nil {
			return err
		}
		if err := buf.Flush(); err != nil {
			return err
		}
	}

	return nil
}

/*
FileDecryptByZyx @Editor robotyang at 2023

# FileDecryptByZyx 文件解密；

@Param srcFilePath：需解密的 文件路径；

@Param dstFilePath：解密后 文件路径；

@Param secret：密钥；

@Reference https://www.jianshu.com/p/0caab60fea9f
*/
func FileDecryptByZyx(srcFilePath, dstFilePath string, secret []byte) (err error) {
	f, err := os.Open(srcFilePath)
	if err != nil { //未找到文件
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	ff, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil { //文件写入错误
		return err
	}
	defer ff.Close()

	//逐行读取密文，进行解密，并且写入文件
	num := 0 //解密次数
	for {
		num = num + 1
		a, err := br.ReadString('\n')
		if err != nil {
			break
		}
		getByte, err := AesBs64Decrypt(a, secret)
		if err != nil { //解密错误
			return err
		}

		buf := bufio.NewWriter(ff)
		if _, err := buf.Write(getByte); err != nil {
			return err
		}
		if err := buf.Flush(); err != nil {
			return err
		}
	}

	return nil
}
