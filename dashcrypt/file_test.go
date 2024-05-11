package dashcrypt_test

import (
	"github.com/rbtyang/godash/dashcrypt"
	"github.com/rbtyang/godash/dashfile"
	"github.com/rbtyang/godash/dashrand"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const filePrefix = "temp/dashcrypt_file_test"
const originFilePath = filePrefix + ".md"

/*
init is a ...
*/
func init() {
	log.Println("Before file_test.go tests")

	// 初始化测试数据
	if !dashfile.IsFile(originFilePath) {
		file, err := dashfile.CreateOrReset(originFilePath)
		if err != nil {
			log.Panicln(err)
			return
		}
		defer file.Close()

		file.WriteString(`
# godash > dashcrypt
> dashcrypt 是 godash 工具集里的 关于加解密的模块

---

## 已支持算法

### AES
  - AES 字符 加解密
    - AES CBC/ECB 加解密
    - AES CBC解密（支持前后端）
  - AES 文件 加解密

---

## 贡献人
- 大绵羊（rbtyang）
- !@#$%^&*())_+
`)
	}
}

/*
TestFileEncrypt is a ...
*/
func TestFileEncrypt(t *testing.T) {
	cryptFilePath := filePrefix + "_crypt.md"
	plainFilePath := filePrefix + "_plain.md"

	// 执行测试用例
	{
		secret := []byte(dashrand.Str(dashrand.ModeNumAlphaSp, 32))

		err := dashcrypt.FileEncryptByZyx(originFilePath, cryptFilePath, secret)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("File Encrypt Success")

		err2 := dashcrypt.FileDecryptByZyx(cryptFilePath, plainFilePath, secret)
		if err2 != nil {
			t.Error(err2)
			return
		}
		t.Log("File Decrypt Success")

		assert.Equal(t, dashfile.CompareFileBySum(originFilePath, plainFilePath), true)
		assert.Equal(t, dashfile.CompareFileBySum(cryptFilePath, plainFilePath), false)
		t.Log("File CompareFileBySum Success")
	}
}
