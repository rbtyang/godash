//只是针对git.myscrm.cn/golang/common/json中的方法加入了错误日志记录
package jsondash

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/rbtyang/godash/convdash"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

func Marshal(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("invalid memory address or nil pointer dereference")
	}
	msByte, err := MarshalToByte(v)
	return convdash.Byte2Str(msByte), err
}

func MarshalNoError(v interface{}) string {
	str, _ := Marshal(v)
	return str
}

func MarshalToByte(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(convdash.Str2ByteByReflect(data), v)
}

func UnmarshalFuzzyDecoders(data string, v interface{}) error {
	return jsoniter.UnmarshalFromString(data, v)
}

func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
