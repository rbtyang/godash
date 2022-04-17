package convdash

import (
	"encoding/json"
	"errors"
	"reflect"
	"unsafe"
)

func StrToByte(s string) []byte {
	return []byte(s)
}

func ByteToStr(b []byte) string {
	return string(b)
}

func StrToByteByUnsafe(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func ByteToStrByUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// converts string to a byte slice without memory allocation.
// Note it may break if string and/or slice header will change
// in the future go versions.
func StrToByteByReflect(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// 转换 数据为Map结构；
// @param.data 支持 struct、map、slice 以及它们的指针类型；
// @return.maps map[string]interface{}、[]interface{}；
func ObjToMap(data interface{}) (maps interface{}, err error) {
	jsonb, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	rType := reflect.TypeOf(data)
	rtKind := rType.Kind()
KindSwitch:
	switch rtKind {
	case reflect.Ptr:
		rtKind = rType.Elem().Kind()
		goto KindSwitch
	case reflect.Map, reflect.Struct:
		var rMaps map[string]interface{}
		err = json.Unmarshal(jsonb, &rMaps)
		maps = rMaps
	case reflect.Slice:
		var rMaps []interface{}
		err = json.Unmarshal(jsonb, &rMaps)
		maps = rMaps
	default:
		return nil, errors.New("Unsupport data type")
	}
	if err != nil {
		return nil, err
	}

	return maps, nil
}

