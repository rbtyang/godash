package dashconv

import (
	"encoding/json"
	"errors"
	"reflect"
)

/*
ObjToMap @Editor robotyang at 2023

# ObjToMap 转换 数据为Map结构

@Param data：支持 struct、map、slice 以及它们的指针类型

@Return maps：可为 map[string]any、[]any
*/
func ObjToMap(data any) (maps any, err error) {
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
		var rMaps map[string]any
		err = json.Unmarshal(jsonb, &rMaps)
		maps = rMaps
	case reflect.Slice:
		var rMaps []any
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
