package convdash

import (
	"encoding/json"
	"errors"
	"reflect"
)

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
