package convdash

import "github.com/goinggo/mapstructure"

/*
WeakMapToStcWithTag 解析map 到结构体（自定义tag 作为 key）

@Editor robotyang at 2023
*/
func WeakMapToStcWithTag(input, output interface{}, tagName string) error {
	cfg := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true, //支持弱类型输入
		TagName:          tagName,
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
