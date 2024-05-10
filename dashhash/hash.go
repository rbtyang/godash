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

/*
Md5 @Editor robotyang at 2023
*/
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum(dashconv.StrToByteByReflect(str)))
}

/*
Md5ByteToStr @Editor robotyang at 2023
*/
func Md5ByteToStr(byts []byte) string {
	return fmt.Sprintf("%x", md5.Sum(byts))
}

/*
Sha1 @Editor robotyang at 2023
*/
func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum(dashconv.StrToByteByReflect(str)))
}

/*
Sha256 @Editor robotyang at 2023
*/
func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256(dashconv.StrToByteByReflect(str)))
}

/*
Sha512 @Editor robotyang at 2023
*/
func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512(dashconv.StrToByteByReflect(str)))
}

/*
Hmac @Editor robotyang at 2023

# Hmac 以一个密钥和一个消息为输入，生成一个消息摘要作为输出。
*/
func Hmac(str, salt string) string {
	hmacIt := hmac.New(md5.New, []byte(salt))
	hmacIt.Write([]byte(str))
	return hex.EncodeToString(hmacIt.Sum([]byte("")))
}
