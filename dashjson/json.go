package dashjson

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/rbtyang/godash/dashconv"
)

/*
init is a ...
*/
func init() {
	extra.RegisterFuzzyDecoders()
}

/*
Marshal is a ...
*/
func Marshal(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("invalid memory address or nil pointer dereference")
	}
	msByte, err := MarshalToByte(v)
	return dashconv.ByteToStrByUnsafe(msByte), err
}

/*
MarshalNoErr is a ...
*/
func MarshalNoErr(v interface{}) string {
	str, _ := Marshal(v)
	return str
}

/*
Unmarshal is a ...
*/
func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(dashconv.StrToByteByReflect(data), v)
}

/*
UnmarshalFuzzyDecoders is a ...
*/
func UnmarshalFuzzyDecoders(data string, v interface{}) error {
	return jsoniter.UnmarshalFromString(data, v)
}

/*
MarshalToByte is a ...
*/
func MarshalToByte(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

/*
UnmarshalByte is a ...
*/
func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
