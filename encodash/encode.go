package encodash

import (
	"encoding/base64"
)

func Base64Encrypt(plainstr string) string {
	cipherstr := base64.StdEncoding.EncodeToString([]byte(plainstr))
	return cipherstr
}

func Base64Decrypt(cipherstr string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

// URL和文件名安全方式 是 标准方式的变体，其输出 用于URL和文件名。
// 因为 +和/字符 是标准Base64字符，但对 URL和文件名 是编码不安全的。
// 因此 变体 会使用 -代替+， _（下划线）代替/ 。
func Base64UrlEncrypt(plainstr string) string {
	cipherstr := base64.URLEncoding.EncodeToString([]byte(plainstr))
	return cipherstr
}

// URL和文件名安全方式 是 标准方式的变体，其输出 用于URL和文件名。
// 因为 +和/字符 是标准Base64字符，但对 URL和文件名 是编码不安全的。
// 因此 变体 会使用 -代替+， _（下划线）代替/ 。
func Base64UrlDecrypt(cipherstr string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}
