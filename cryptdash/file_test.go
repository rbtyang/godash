package cryptdash_test

import (
	"github.com/rbtyang/godash/cryptdash"
	"github.com/rbtyang/godash/filedash"
	"github.com/rbtyang/godash/randdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const filePrefix = "temp/cryptdash_file_test"
const originFilePath = filePrefix + ".md"

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before file_test.go tests")

	// 初始化测试数据
	if !filedash.IsExistFile(originFilePath) {
		file, err := filedash.Rebuild(originFilePath)
		if err != nil {
			log.Panicln(err)
			return
		}
		defer file.Close()

		file.WriteString(`
# godash > cryptdash
> cryptdash 是 godash 工具集里的 关于加解密的模块

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

@Editor robotyang at 2023
*/
func TestFileEncrypt(t *testing.T) {
	cryptFilePath := filePrefix + "_crypt.md"
	plainFilePath := filePrefix + "_plain.md"

	// 执行测试用例
	{
		secret := []byte(randdash.Str(randdash.ModeNumAlphaSp, 32))

		err := cryptdash.FileEncryptByZyx(originFilePath, cryptFilePath, secret)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("File Encrypt Success")

		err2 := cryptdash.FileDecryptByZyx(cryptFilePath, plainFilePath, secret)
		if err2 != nil {
			t.Error(err2)
			return
		}
		t.Log("File Decrypt Success")

		assert.Equal(t, filedash.CompareFileBySum(originFilePath, plainFilePath), true)
		assert.Equal(t, filedash.CompareFileBySum(cryptFilePath, plainFilePath), false)
		t.Log("File CompareFileBySum Success")
	}
}
