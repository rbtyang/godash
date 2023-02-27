package convdash

import (
	"reflect"
	"unsafe"
)

/*
StrToByte is a ...

@Editor robotyang at 2023
*/
func StrToByte(s string) []byte {
	return []byte(s)
}

/*
ByteToStr is a ...

@Editor robotyang at 2023
*/
func ByteToStr(b []byte) string {
	return string(b)
}

/*
StrToByteByUnsafe is a ...

@Editor robotyang at 2023
*/
func StrToByteByUnsafe(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/*
ByteToStrByUnsafe is a ...

@Editor robotyang at 2023
*/
func ByteToStrByUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StrToByteByReflect converts string to a byte slice without memory allocation.

/*
StrToByteByReflect  Note it may break if string and/or slice header will change in the future go versions.

@Editor robotyang at 2023
*/
func StrToByteByReflect(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
