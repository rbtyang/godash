package dashjson_test

import (
	"encoding/json"
	"github.com/rbtyang/godash/dashjson"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Student struct {
	Uuid      string   `comment:"学生uuid" remark:"【必填】"`
	Type      string   `comment:"年段" remark:"【必填】如 MIDDLE 初中、HIGH 高中"`
	Name      string   `comment:"姓名" remark:"【必填】"`
	Level     int32    `comment:"评级" remark:"【必填】0~10级"`
	Gender    int32    `comment:"性别" remark:"【可选】0男，1女，默认0"`
	Mobile    string   `comment:"号码" remark:"【可选】年段为高中生时，则联系号码必填"`
	ClassUuid string   `comment:"班级uuid" remark:"【可选】"`
	ClassName string   `comment:"班级名称" remark:"【可选】班级uuid不为空时，则班级名称必填"`
	Hobbies   []string `comment:"兴趣" remark:"【可选】1唱 2跳 3rap 4篮球"`
	Address   *Address `comment:"住所地址" remark:"【可选】默认空"`
	Remark    string   `comment:"备注" remark:"【可选】默认空"`
}
type Address struct {
	Name   string `comment:"地址别名" remark:"【可选】默认空"`
	Detail string `comment:"详细地址" remark:"【可选】默认空"`
}

var student = Student{
	Uuid:      "a223db68-0ea4-11ef-8f30-fa163e930d3e",
	Type:      "MIDDLE",
	Name:      "张三",
	Gender:    1,
	Mobile:    "1234567890",
	Level:     9,
	ClassUuid: "a223db68-0ea4-11ef-8f30",
	ClassName: "三年一班",
	Hobbies:   []string{"篮球", "足球", "排球"},
	Address: &Address{
		Name:   "三里屯",
		Detail: "北京市朝阳区三里屯",
	},
	Remark: "三好学生",
}

/*
@Editor robotyang at 2024

TestEmpty is a ...
*/
func TestMarshal(t *testing.T) {
	cases := []any{
		``,
		`a`,
		student,
	}

	student2 := student
	student2.Hobbies = nil
	student2.Address = nil
	cases = append(cases, student2)

	for _, cs := range cases {
		byt, err := json.Marshal(cs)
		assert.Equal(t, nil, err)
		want := string(byt)

		recv, err := dashjson.Marshal(cs)
		assert.Equal(t, nil, err)

		assert.Equal(t, want, recv)
	}
}

func TestMarshalWithTag(t *testing.T) {
	//TODO here
}
