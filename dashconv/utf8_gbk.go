package dashconv

import (
	"bytes"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

/*
GbkToUtf8 @Editor robotyang at 2023

# GbkToUtf8 转换 GBK字符串 为 UTF8字符串
*/
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

/*
Utf8ToGbk @Editor robotyang at 2023

# Utf8ToGbk 转换 UTF8字符串 为 GBK字符串
*/
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

/*
TransEncoding @Editor robotyang at 2023

# TransEncoding 转换 字符串 编码类型

@Param src：原字符串

@Param srcCode：原字符串 编码类型，如 "gbk"、"utf-8"

@Param dstCode：目标字符串 编码类型，如 "gbk"、"utf-8"
*/
func TransEncoding(src string, srcCode string, dstCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(dstCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)

	return result
}
