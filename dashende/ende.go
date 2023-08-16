package dashende

import (
	"encoding/base64"
)

/*
Base64Encode is a ...
*/
func Base64Encode(plainstr string) string {
	cipherstr := base64.StdEncoding.EncodeToString([]byte(plainstr))
	return cipherstr
}

/*
Base64Decode is a ...
*/
func Base64Decode(cipherstr string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

/*
Base64UrlEncode is a ...

URL和文件名安全方式 是 标准方式的变体，其输出 用于URL和文件名。

因为 +和/字符 是标准Base64字符，但对 URL和文件名 是编码不安全的。

因此 变体 会使用 -代替+， _（下划线）代替/ 。
*/
func Base64UrlEncode(plainstr string) string {
	cipherstr := base64.URLEncoding.EncodeToString([]byte(plainstr))
	return cipherstr
}

/*
Base64UrlDecode is a ...

URL和文件名安全方式 是 标准方式的变体，其输出 用于URL和文件名。

因为 +和/字符 是标准Base64字符，但对 URL和文件名 是编码不安全的。

因此 变体 会使用 -代替+， _（下划线）代替/ 。
*/
func Base64UrlDecode(cipherstr string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}
