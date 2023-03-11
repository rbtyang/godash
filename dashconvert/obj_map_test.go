package dashconvert_test

import (
	"github.com/rbtyang/godash/dashconvert"
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

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before this tests")
}

/*
TestObjToMap is a ...

@Editor robotyang at 2023
*/
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
		orderMap, err := dashconvert.ObjToMap(order)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("orderMap: ", orderMap)
	}
	{ // map
		orderList := map[string]*Order{"aaa": order, "bbb": order, "ccc": nil}
		orderMap, err := dashconvert.ObjToMap(orderList)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("orderMap: ", orderMap)
	}
	{ // slice
		orderList := []*Order{order, order, nil}
		orderMap, err := dashconvert.ObjToMap(orderList)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("orderMap: ", orderMap)
	}
}
