package dashconv

import "github.com/goinggo/mapstructure"

/*
WeakMapToStructWithTag @Editor robotyang at 2023

# WeakMapToStructWithTag 解析map到结构体（自定义tag 作为 key）

@Param inMap：输入的map

@Param outStruct：接收的结构体

@Param tagName：结构体的关联tag名
*/
func WeakMapToStructWithTag(inMap, outStruct any, tagName string) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           outStruct,
		WeaklyTypedInput: true,    //支持弱类型输入
		TagName:          tagName, //defaults to "mapstructure"
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return err
	}
	return decoder.Decode(inMap)
}
