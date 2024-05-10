package dashvalid

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type Valid struct {
	separate string //消息分隔符
	*validator.Validate
	trans ut.Translator
}

/*
New @Editor robotyang at 2023

# New 支持 string msg, err obj

@Reference https://juejin.cn/post/6847902214279659533
*/
func New() *Valid {
	v := &Valid{
		separate: "; ",
		trans:    nil,
		Validate: validator.New(),
	}

	//注册自定义标签
	v.Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		comment := field.Tag.Get("comment")
		if comment == "" {
			return field.Name
		}
		return comment
	})

	//注册中文翻译器
	zhs := zh.New()
	uni := ut.New(zhs, zhs)
	trans, _ := uni.GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(v.Validate, trans)
	v.trans = trans

	return v
}

// Sep 设置 消息分隔符
func (v *Valid) Sep(connector string) *Valid {
	v.separate = connector
	return v
}

// List 返回 消息切片（不带分隔符）
func (v *Valid) List(err error) []string {
	var errList []string
	errs := err.(validator.ValidationErrors)
	for _, e := range errs { // can translate each error one at a time.
		errList = append(errList, e.Translate(v.trans))
	}
	return errList
}

// Str 返回 消息字符串（分隔符区分）
func (v *Valid) Str(err error) string {
	if err == nil {
		return ""
	} else {
		return strings.Join(v.List(err), v.separate)
	}
}
