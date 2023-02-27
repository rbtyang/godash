//只是针对git.myscrm.cn/golang/common/json中的方法加入了错误日志记录
package jsondash

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/rbtyang/godash/convdash"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	extra.RegisterFuzzyDecoders()
}

/*
Marshal is a ...

@Editor robotyang at 2023
*/
func Marshal(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("invalid memory address or nil pointer dereference")
	}
	msByte, err := MarshalToByte(v)
	return convdash.ByteToStrByUnsafe(msByte), err
}

/*
MarshalNoError is a ...

@Editor robotyang at 2023
*/
func MarshalNoError(v interface{}) string {
	str, _ := Marshal(v)
	return str
}

/*
MarshalToByte is a ...

@Editor robotyang at 2023
*/
func MarshalToByte(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

/*
Unmarshal is a ...

@Editor robotyang at 2023
*/
func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(convdash.StrToByteByReflect(data), v)
}

/*
UnmarshalFuzzyDecoders is a ...

@Editor robotyang at 2023
*/
func UnmarshalFuzzyDecoders(data string, v interface{}) error {
	return jsoniter.UnmarshalFromString(data, v)
}

/*
UnmarshalByte is a ...

@Editor robotyang at 2023
*/
func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
