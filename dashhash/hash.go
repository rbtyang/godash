package dashhash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/rbtyang/godash/dashconv"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum(dashconv.StrToByteByReflect(str)))
}

func Md5ByteToStr(byts []byte) string {
	return fmt.Sprintf("%x", md5.Sum(byts))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum(dashconv.StrToByteByReflect(str)))
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256(dashconv.StrToByteByReflect(str)))
}

func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512(dashconv.StrToByteByReflect(str)))
}

/*
Hmac 以一个密钥和一个消息为输入，生成一个消息摘要作为输出。
*/
func Hmac(salt, str string) string {
	hmacIt := hmac.New(md5.New, []byte(salt))
	hmacIt.Write([]byte(str))
	return hex.EncodeToString(hmacIt.Sum([]byte("")))
}
