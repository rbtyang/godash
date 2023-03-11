package dashcrypt_test

import (
	"github.com/rbtyang/godash/dashcrypt"
	"log"
	"testing"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before sign_test.go tests")
}

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
TestSignFlatMap is a ...

@Editor robotyang at 2023
*/
func TestSignFlatMap(t *testing.T) {
	secret := "9adceaa7c1c1979214becf7747e05daa"
	data := map[string]string{
		"timestamp": "1554208460",
		"abcc":      "aaa",
		"cccc":      "bbbbb",
		"xxx":       "asdf123",
		"sign":      "C8FB8137FF68345F04CE8BE6A1F84924",
	}

	sign := dashcrypt.SignFlatMap(data, secret)
	t.Log("SignFlatMap: ", sign)

	ckRes := dashcrypt.CheckSignFlatMap(data, secret, sign)
	t.Log("CheckSignFlatMap: ", ckRes)
}

// //TODO 支持各种结构的签名计算
//func TestSignObj(t *testing.T) {
//	secret := "9adceaa7c1c1979214becf7747e05daa"
//	order := &Order{
//		Id:    111,
//		Name:  "asd哈哈",
//		Label: []string{"asd", "哈哈", "asd哈哈"},
//		Address: &Address{
//			City:   "北京",
//			Street: "朝阳街",
//		},
//		Items: []*Item{
//			{
//				Product: "苹果",
//				Number:  123,
//				Stock: &Stock{
//					Place:    "3区4库",
//					Quantity: 1234,
//				},
//			},
//			{
//				Product: "香蕉",
//				Number:  456,
//				Stock: &Stock{
//					Quantity: 4567,
//					Place:    "5区3库",
//				},
//			},
//		},
//	}
//
//	//data := order
//	//data := map[string]*Order{"aaa": order, "bbb": order, "ccc": nil}
//	data := []*Order{order, order, nil}
//
//	sign, err := cryptdash.SignObj(data, secret)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	t.Log("SignObj: ", sign)
//
//	ckRes := cryptdash.CheckSignObj(data, secret, sign)
//	t.Log("CheckSignObj: ", ckRes)
//}
//
//func TestMakeSignByUrlParam(t *testing.T) {
//	queryMap := map[string]string{
//		"abcc": "aaa",
//		"cccc": "bbbbb",
//		"xxx":  "asdf123",
//	}
//	queryParams := url.Values{}
//	for k, v := range queryMap {
//		queryParams.Add(k, v)
//	}
//	queryStr := queryParams.Encode()
//	t.Log(queryStr)
//
//	bodyParams := url.Values{}
//	bodyParams.Add("arrtest", "aaa")
//	bodyParams.Add("arrtest", "bbb")
//	bodyParams.Add("arrtest", "哈哈")
//	bodyParams.Add("aaatest", "嘿嘿")
//	bodyStr := bodyParams.Encode()
//	t.Log(bodyStr)
//}
