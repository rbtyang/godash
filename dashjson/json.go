package dashjson

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/rbtyang/godash/dashconv"
)

/*
init @Editor robotyang at 2023

# init is a ...
*/
func init() {
	extra.RegisterFuzzyDecoders()
}

/*
Marshal @Editor robotyang at 2023

# Marshal is a ...
*/
func Marshal(v any) (string, error) {
	if v == nil {
		return "", errors.New("invalid memory address or nil pointer dereference")
	}
	msByte, err := MarshalToByte(v)
	return dashconv.ByteToStrByUnsafe(msByte), err
}

/*
MarshalNoErr @Editor robotyang at 2023

# MarshalNoErr is a ...
*/
func MarshalNoErr(v any) string {
	str, _ := Marshal(v)
	return str
}

/*
Unmarshal @Editor robotyang at 2023

# Unmarshal is a ...
*/
func Unmarshal(data string, v any) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(dashconv.StrToByteByReflect(data), v)
}

/*
UnmarshalFuzzyDecoders @Editor robotyang at 2023

# UnmarshalFuzzyDecoders is a ...
*/
func UnmarshalFuzzyDecoders(data string, v any) error {
	return jsoniter.UnmarshalFromString(data, v)
}

/*
MarshalToByte @Editor robotyang at 2023

# MarshalToByte is a ...
*/
func MarshalToByte(v any) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

/*
UnmarshalByte @Editor robotyang at 2023

# UnmarshalByte is a ...
*/
func UnmarshalByte(data []byte, v any) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
