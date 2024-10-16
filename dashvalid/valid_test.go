package dashvalid_test

import (
	"testing"

	"github.com/rbtyang/godash/dashvalid"
	"github.com/stretchr/testify/assert"
)

var v = dashvalid.New()

/*
验证器可参考：@Reference https://juejin.cn/post/6847902214279659533
*/
type Student struct {
	Uuid      string   `validate:"required,len=36" comment:"学生uuid" remark:"【必填】"`
	Type      string   `validate:"required,oneof=MIDDLE HIGH" comment:"年段" remark:"【必填】如 MIDDLE 初中、HIGH 高中"`
	Name      string   `validate:"required,max=5" comment:"姓名" remark:"【必填】"`
	Level     int32    `validate:"number,oneof=0 1 2 3 4 5 6 7 8 9 10" comment:"评级" remark:"【必填】0~10级"`
	Gender    int32    `validate:"oneof=0 1" comment:"性别" remark:"【可选】0男，1女，默认0"`
	Mobile    string   `validate:"required_if=Type HIGH,max=11" comment:"号码" remark:"【可选】年段为高中生时，则联系号码必填"`
	ClassUuid string   `validate:"max=36" comment:"班级uuid" remark:"【可选】"`
	ClassName string   `validate:"required_with=ClassUuid,max=100" comment:"班级名称" remark:"【可选】班级uuid不为空时，则班级名称必填"`
	Hobbies   []string `validate:"dive,oneof=1 2 3 4" comment:"兴趣" remark:"【可选】1唱 2跳 3rap 4篮球"`
	Remark    string   `validate:"" comment:"备注" remark:"【可选】默认空"`
}

var student = Student{
	Uuid:      "a223db68-0ea4-11ef-8f30-fa163e930d3e",
	Type:      "MIDDLE",
	Name:      "张三",
	Gender:    1,
	Mobile:    "",
	Level:     9,
	ClassUuid: "",
	ClassName: "",
	Hobbies:   nil,
	Remark:    "",
}

/*
@Editor robotyang at 2024

TestValid is a ...
*/
func TestValid(t *testing.T) {
	{
		dto := student
		dto.Uuid = ""
		want := "学生uuid为必填字段"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
	{
		dto := student
		dto.Uuid = "abc"
		want := "学生uuid长度必须是36个字符"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
	{
		dto := student
		dto.Type = "XXX"
		want := "年段必须是[MIDDLE HIGH]中的一个"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
	{
		dto := student
		dto.Name = "abcdefg"
		want := "姓名长度不能超过5个字符"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
	{
		dto := student
		dto.Type = "MIDDLE"
		dto.Mobile = ""
		want := ""
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
	{
		dto := student
		dto.Type = "HIGH"
		dto.Mobile = ""
		want := "号码为必填字段"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
	{
		dto := student
		dto.ClassUuid = "a66e4df2-0ea4-11ef-8f30-fa163e930d3e"
		dto.ClassName = ""
		want := "班级名称为必填字段"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}

	{
		dto := student
		dto.Uuid = ""
		dto.Type = ""
		want := "学生uuid为必填字段; 年段为必填字段"
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}

	{
		dto := student
		dto.Hobbies = []string{"8", "1", "8"}
		want := "兴趣[0]必须是[1 2 3 4]中的一个; 兴趣[2]必须是[1 2 3 4]中的一个" //
		recv := v.Str(v.Struct(dto))
		assert.Equal(t, want, recv)
	}
}
