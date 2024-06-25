package dashstr_test

import (
	"github.com/rbtyang/godash/dashstr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsZHMobile(t *testing.T) {
	inputMap := map[string]bool{
		"12345678910":     true,  //大陆手机
		"8612345678910":   true,  //大陆手机
		"86-12345678910":  true,  //大陆手机
		"+8612345678910":  true,  //大陆手机
		"+86-12345678910": true,  //大陆手机
		"13912345678":     true,  //大陆手机
		"8613912345678":   true,  //大陆手机
		"86-13912345678":  true,  //大陆手机
		"+8613912345678":  true,  //大陆手机
		"+86-13912345678": true,  //大陆手机
		"00852123456":     false, //港号码（含区号）
		"0085212345678":   false, //港号码（含区号）
		"00852123456789":  false, //港号码（含区号）
		"00853123456":     false, //澳号码（含区号）
		"008531234567":    false, //澳号码（含区号）
		"00853123456789":  false, //澳号码（含区号）
		"00886123456":     false, //台号码（含区号）
		"008861234567":    false, //台号码（含区号）
		"00886123456789":  false, //台号码（含区号）
		"1234567":         false, //大陆座机
		"12345678":        false, //大陆座机
		"86287788":        false, //大陆座机
		"089886287788":    false, //大陆座机（含区号）
		"0898-86287788":   false, //大陆座机（含区号）
		"1":               false, //其他号码
		"12":              false, //其他号码
		"123":             false, //其他号码
		"234":             false, //其他号码
		"123456789":       false, //其他号码
		"12345678910a":    false, //其他号码
		"1234567891011":   false, //其他号码
		"0898-862877889":  false, //其他号码
	}
	for tel, want := range inputMap {
		recv := dashstr.IsZHMobile(tel)
		assert.Equal(t, want, recv)
	}
}

func Test_IsZHLandline(t *testing.T) {
	inputMap := map[string]bool{
		"12345678910":     false, //大陆手机
		"8612345678910":   false, //大陆手机
		"86-12345678910":  false, //大陆手机
		"+8612345678910":  false, //大陆手机
		"+86-12345678910": false, //大陆手机
		"13912345678":     false, //大陆手机
		"8613912345678":   false, //大陆手机
		"86-13912345678":  false, //大陆手机
		"+8613912345678":  false, //大陆手机
		"+86-13912345678": false, //大陆手机
		"00852123456":     false, //港号码（含区号）
		"0085212345678":   false, //港号码（含区号）
		"00852123456789":  false, //港号码（含区号）
		"00853123456":     false, //澳号码（含区号）
		"008531234567":    false, //澳号码（含区号）
		"00853123456789":  false, //澳号码（含区号）
		"00886123456":     false, //台号码（含区号）
		"008861234567":    false, //台号码（含区号）
		"00886123456789":  false, //台号码（含区号）
		"1234567":         true,  //大陆座机
		"12345678":        true,  //大陆座机
		"86287788":        true,  //大陆座机
		"089886287788":    true,  //大陆座机（含区号）
		"0898-86287788":   true,  //大陆座机（含区号）
		"1":               false, //其他号码
		"12":              false, //其他号码
		"123":             false, //其他号码
		"234":             false, //其他号码
		"123456789":       false, //其他号码
		"12345678910a":    false, //其他号码
		"1234567891011":   false, //其他号码
		"0898-862877889":  false, //其他号码
	}
	for tel, want := range inputMap {
		recv := dashstr.IsZHLandline(tel)
		bol := assert.Equal(t, want, recv)
		if !bol {
			println(tel)
		}
	}
}

func Test_ParseTelType(t *testing.T) {
	inputMap := map[string]dashstr.TelType{
		"12345678910":     dashstr.TelType_ZH,   //大陆手机
		"8612345678910":   dashstr.TelType_ZH,   //大陆手机
		"86-12345678910":  dashstr.TelType_ZH,   //大陆手机
		"+8612345678910":  dashstr.TelType_ZH,   //大陆手机
		"+86-12345678910": dashstr.TelType_ZH,   //大陆手机
		"13912345678":     dashstr.TelType_ZH,   //大陆手机
		"8613912345678":   dashstr.TelType_ZH,   //大陆手机
		"86-13912345678":  dashstr.TelType_ZH,   //大陆手机
		"+8613912345678":  dashstr.TelType_ZH,   //大陆手机
		"+86-13912345678": dashstr.TelType_ZH,   //大陆手机
		"00852123456":     dashstr.TelType_HK,   //港号码（含区号）
		"0085212345678":   dashstr.TelType_HK,   //港号码（含区号）
		"00852123456789":  dashstr.TelType_HK,   //港号码（含区号）
		"00853123456":     dashstr.TelType_AM,   //澳号码（含区号）
		"008531234567":    dashstr.TelType_AM,   //澳号码（含区号）
		"00853123456789":  dashstr.TelType_AM,   //澳号码（含区号）
		"00886123456":     dashstr.TelType_TW,   //台号码（含区号）
		"008861234567":    dashstr.TelType_TW,   //台号码（含区号）
		"00886123456789":  dashstr.TelType_TW,   //台号码（含区号）
		"1234567":         dashstr.TelType_ZHLD, //大陆座机
		"12345678":        dashstr.TelType_ZHLD, //大陆座机
		"86287788":        dashstr.TelType_ZHLD, //大陆座机
		"089886287788":    dashstr.TelType_ZHLD, //大陆座机（含区号）
		"0898-86287788":   dashstr.TelType_ZHLD, //大陆座机（含区号）
		"1":               dashstr.TelType_QT,   //其他号码
		"12":              dashstr.TelType_QT,   //其他号码
		"123":             dashstr.TelType_QT,   //其他号码
		"234":             dashstr.TelType_QT,   //其他号码
		"123456789":       dashstr.TelType_QT,   //其他号码
		"12345678910a":    dashstr.TelType_QT,   //其他号码
		"1234567891011":   dashstr.TelType_QT,   //其他号码
		"0898-862877889":  dashstr.TelType_QT,   //其他号码
	}
	for tel, want := range inputMap {
		recv := dashstr.ParseTelType(tel)
		assert.Equal(t, want, recv)
	}
}
