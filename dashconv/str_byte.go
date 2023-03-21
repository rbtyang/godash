package dashconv

import (
	"reflect"
	"unsafe"
)

/*
@Editor robotyang at 2023

StrToByte is a ...
*/
func StrToByte(s string) []byte {
	return []byte(s)
}

/*
@Editor robotyang at 2023

ByteToStr is a ...
*/
func ByteToStr(b []byte) string {
	return string(b)
}

/*
@Editor robotyang at 2023

StrToByteByUnsafe is a ...
*/
func StrToByteByUnsafe(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/*
@Editor robotyang at 2023

ByteToStrByUnsafe is a ...
*/
func ByteToStrByUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

/*
@Editor robotyang at 2023

StrToByteByReflect  converts string to a byte slice without memory allocation.
Note it may break if string and/or slice header will change in the future go versions.
*/
func StrToByteByReflect(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
