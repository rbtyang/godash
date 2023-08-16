package dashconv

import "github.com/goinggo/mapstructure"

/*
WeakMapToStructWithTag 解析map 到结构体（自定义tag 作为 key）
*/
func WeakMapToStructWithTag(inMap, outStruct interface{}, tagName string) error {
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
