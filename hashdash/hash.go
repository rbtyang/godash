package hashdash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/rbtyang/godash/convdash"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum(convdash.StrToByteByReflect(str)))
}

func Md5ByteToStr(byts []byte) string {
	return fmt.Sprintf("%x", md5.Sum(byts))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum(convdash.StrToByteByReflect(str)))
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256(convdash.StrToByteByReflect(str)))
}

func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512(convdash.StrToByteByReflect(str)))
}
