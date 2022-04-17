// @reference https://www.jianshu.com/p/0caab60fea9f
package cryptdash

import (
	"bufio"
	"github.com/rbtyang/godash/logdash"
	"os"
)

//EncryptFile 文件加密，filePath 需要加密的文件路径 ，newName 加密后文件名
func FileEncryptByZyx(srcFilePath, dstFilePath string, secret []byte) (err error) {
	clsPre, _ := logdash.Pre("FileEncryptByZyx")
	defer clsPre()

	f, err := os.Open(srcFilePath)
	if err != nil {
		logdash.Error("未找到文件")
		return
	}
	defer f.Close()

	fInfo, _ := f.Stat()
	fLen := fInfo.Size()
	logdash.Info("待处理文件大小：", fLen)

	maxLen := 1024 * 1024 * 100 //100mb  每 100mb 进行加密一次
	var forNum int64 = 0
	getLen := fLen

	if fLen > int64(maxLen) {
		getLen = int64(maxLen)
		forNum = fLen / int64(maxLen)
		logdash.Info("需要加密次数：", forNum+1)
	}

	//加密后存储的文件
	ff, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logdash.Error("文件写入错误")
		return err
	}
	defer ff.Close()

	//循环加密，并且写入文件
	for i := 0; i < int(forNum+1); i++ {
		a := make([]byte, getLen)
		n, err := f.Read(a)
		if err != nil {
			logdash.Error("文件读取错误")
			return err
		}
		getByte, err := AesEncrypt(a[:n], secret)
		if err != nil {
			logdash.Error("加密错误")
			return err
		}
		//换行处理，有点乱了，想到更好的再改
		getBytes := append([]byte(getByte), []byte("\n")...)
		//写入
		buf := bufio.NewWriter(ff)
		buf.WriteString(string(getBytes[:]))
		buf.Flush()
	}

	ffInfo, _ := ff.Stat()
	logdash.Infof("文件加密成功，生成文件名为：%s，文件大小为：%v Byte \n", ffInfo.Name(), ffInfo.Size())

	return nil
}

//EncryptFile 文件解密，filePath 需要解密的文件路径 ，newName 解密后文件名
func FileDecryptByZyx(srcFilePath, dstFilePath string, secret []byte) (err error) {
	clsPre, _ := logdash.Pre("FileDecryptByZyx")
	defer clsPre()

	f, err := os.Open(srcFilePath)
	if err != nil {
		logdash.Error("未找到文件")
		return
	}
	defer f.Close()

	fInfo, _ := f.Stat()
	logdash.Info("待处理文件大小:", fInfo.Size())

	br := bufio.NewReader(f)
	ff, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logdash.Error("文件写入错误")
		return err
	}
	defer ff.Close()

	//逐行读取密文，进行解密，并且写入文件
	num := 0
	for {
		num = num + 1
		a, err := br.ReadString('\n')
		if err != nil {
			break
		}
		getByte, err := AesDecrypt(a, secret)
		if err != nil {
			logdash.Error("解密错误")
			return err
		}

		buf := bufio.NewWriter(ff)
		buf.Write(getByte)
		buf.Flush()
	}
	logdash.Info("解密次数：", num)

	ffInfo, _ := ff.Stat()
	logdash.Infof("文件解密成功，生成文件名为：%s，文件大小为：%v Byte \n", ffInfo.Name(), ffInfo.Size())

	return nil
}
