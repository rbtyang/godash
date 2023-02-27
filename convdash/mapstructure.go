package convdash

import "github.com/goinggo/mapstructure"

/*
WeakMapToStcWithTag 解析map 到结构体（自定义tag 作为 key）

@Editor robotyang at 2023
*/
func WeakMapToStcWithTag(inMap, outStruct interface{}, tagName string) error {
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
