package dashconv

import (
	"reflect"
	"unsafe"
)

/*
StrToByte @Editor robotyang at 2023

# StrToByte 字符串 转 字节
*/
func StrToByte(s string) []byte {
	return []byte(s)
}

/*
ByteToStr @Editor robotyang at 2023

# ByteToStr 字节 转 字符串
*/
func ByteToStr(b []byte) string {
	return string(b)
}

/*
StrToByteByUnsafe @Editor robotyang at 2023

# StrToByteByUnsafe 字符串 转 字节
*/
func StrToByteByUnsafe(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/*
ByteToStrByUnsafe @Editor robotyang at 2023

# ByteToStrByUnsafe 字节 转 字符串
*/
func ByteToStrByUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

/*
StrToByteByReflect @Editor robotyang at 2023

# StrToByteByReflect 将字符串转换为字节片而不分配内存。注意，如果字符串和/或切片头在未来的go版本中改变，它可能会中断。

# StrToByteByReflect converts string to a byte slice without memory allocation.
Note it may break if string and/or slice header will change in the future go versions.
*/
func StrToByteByReflect(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
