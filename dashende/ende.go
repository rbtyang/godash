package dashende

import (
	"encoding/base64"
)

/*
Base64Encode @Editor robotyang at 2023

# Base64Encode 将 明文字符串 编码为 base64字符串（StdEncoding）
*/
func Base64Encode(plainstr string) string {
	cipherstr := base64.StdEncoding.EncodeToString([]byte(plainstr))
	return cipherstr
}

/*
Base64Decode @Editor robotyang at 2023

# Base64Decode 将 base64字符串 解码为 明文字符串（StdEncoding）
*/
func Base64Decode(cipherstr string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

/*
Base64UrlEncode @Editor robotyang at 2023

# Base64UrlEncode 将 明文字符串 编码为 base64字符串（URLEncoding，它通常用于url和文件名中。）

URL和文件名安全方式 是 标准方式的变体，其输出 用于URL和文件名。

因为 +和/字符 是标准Base64字符，但对 URL和文件名 是编码不安全的。

因此 变体 会使用 -代替+， _（下划线）代替/ 。
*/
func Base64UrlEncode(plainstr string) string {
	cipherstr := base64.URLEncoding.EncodeToString([]byte(plainstr))
	return cipherstr
}

/*
Base64UrlDecode @Editor robotyang at 2023

# Base64UrlDecode 将 base64字符串 解码为 明文字符串（URLEncoding，它通常用于url和文件名中。）

@Remark URL和文件名安全方式 是 标准方式的变体，其输出 用于URL和文件名。因为 +和/字符 是标准Base64字符，但对 URL和文件名 是编码不安全的。因此 变体 会使用 -代替+， _（下划线）代替/ 。
*/
func Base64UrlDecode(cipherstr string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}
