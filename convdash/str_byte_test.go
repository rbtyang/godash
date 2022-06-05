package convdash_test

import (
	"github.com/rbtyang/godash/convdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Order struct {
	Id      int
	Name    string
	Label   []string
	Address *Address
	Items   []*Item
}
type Address struct {
	City   string
	Street string
}
type Item struct {
	Product string
	Number  int
	Stock   *Stock
}
type Stock struct {
	Place    string
	Quantity int
}

func init() {
	log.Println("Before this tests")
}

func TestStrToByte(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := convdash.StrToByte("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := "hello world 123 哈哈"
		recv := convdash.ByteToStr([]byte("hello world 123 哈哈"))
		assert.Equal(t, want, recv)
	}
}

func TestStrToByteByUnsafe(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := convdash.StrToByteByUnsafe("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
	{
		want := "hello world 123 哈哈"
		recv := convdash.ByteToStrByUnsafe([]byte("hello world 123 哈哈"))
		assert.Equal(t, want, recv)
	}
}

func TestStrToByteByReflect(t *testing.T) {
	{
		want := []byte("hello world 123 哈哈")
		recv := convdash.StrToByteByReflect("hello world 123 哈哈")
		assert.Equal(t, want, recv)
	}
}

func TestObjToMap(t *testing.T) {
	order := &Order{
		Id:    111,
		Name:  "asd哈哈",
		Label: []string{"asd", "哈哈", "asd哈哈"},
		Address: &Address{
			City:   "北京",
			Street: "朝阳街",
		},
		Items: []*Item{
			{
				Product: "苹果",
				Number:  123,
				Stock: &Stock{
					Place:    "3区4库",
					Quantity: 1234,
				},
			},
			{
				Product: "香蕉",
				Number:  456,
				Stock: &Stock{
					Quantity: 4567,
					Place:    "5区3库",
				},
			},
		},
	}

	{ // struct
		orderMap, err := convdash.ObjToMap(order)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("orderMap: ", orderMap)
	}
	{ // map
		orderList := map[string]*Order{"aaa": order, "bbb": order, "ccc": nil}
		orderMap, err := convdash.ObjToMap(orderList)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("orderMap: ", orderMap)
	}
	{ // slice
		orderList := []*Order{order, order, nil}
		orderMap, err := convdash.ObjToMap(orderList)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("orderMap: ", orderMap)
	}
}
