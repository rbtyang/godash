package dashconv

import (
	"reflect"
	"unsafe"
)

/*
StrToByte is a ...
*/
func StrToByte(s string) []byte {
	return []byte(s)
}

/*
ByteToStr is a ...
*/
func ByteToStr(b []byte) string {
	return string(b)
}

/*
StrToByteByUnsafe is a ...
*/
func StrToByteByUnsafe(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/*
ByteToStrByUnsafe is a ...
*/
func ByteToStrByUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

/*
StrToByteByReflect  converts string to a byte slice without memory allocation.
Note it may break if string and/or slice header will change in the future go versions.
*/
func StrToByteByReflect(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
