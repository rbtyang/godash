package dashcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

/*
AesDecrypt CBC模式解密（支持前后端，是 AesDecryptJsHex 的别名）；

@Param cipherstr 密文（js的生成的密文后 进行了16进制的 hex.encoding，因此在调用该方法之前 go必须要进行 hex.DecodeString）；

@Param secret 密钥；

@Editor robotyang at 2023
*/
func AesDecrypt(cipherstr, secret string) (string, error) {
	return AesDecryptJsHex(cipherstr, secret)
}

//-------------------------------------------------------------------------------------------------------

/*
AesDecryptJsHex CBC模式解密（支持前后端，先16进制解码，再走底层方法 AesDecryptJs 解密）；

@Param cipherstr 密文（js的生成的密文后 进行了16进制的 hex.encoding，因此在调用该方法之前 go必须要进行 hex.DecodeString）；

@Param secret 密钥；

@Reference https://mojotv.cn/go/crypto-js-with-golang

@Editor robotyang at 2023
*/
func AesDecryptJsHex(cipherstr, secret string) (string, error) {
	ciphertext, err := hex.DecodeString(cipherstr)
	if err != nil {
		return "", err
	}
	str, err := AesDecryptJs(ciphertext, []byte(secret))
	if err != nil {
		return "", err
	}
	return string(str), nil
}

//-------------------------------------------------------------------------------------------------------

/*
AesDecryptJs CBC模式解密（支持前后端）；

@Param ciphertext 密文（js的生成的密文后 进行了16进制的 hex.encoding，因此在调用该方法之前 go必须要进行 hex.DecodeString）；

@Param secret 密钥；

@Reference https://mojotv.cn/go/crypto-js-with-golang

@Editor robotyang at 2023
*/
func AesDecryptJs(ciphertext, secret []byte) ([]byte, error) {
	pkey := secretFill(secret, '0', 16) //和js的key补码方法一致
	block, err := aes.NewCipher(pkey)   //选择加密算法
	if err != nil {
		return nil, fmt.Errorf("secret 长度必须 16/24/32长度: %s", err)
	}
	blockModel := cipher.NewCBCDecrypter(block, pkey) //和前端代码对应:   mode: CryptoJS.mode.CBC, CBC算法
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = pKCS7UnFill(plantText) //和前端代码对应:  padding: CryptoJS.pad.Pkcs7
	return plantText, nil
}

/*
secretFill 对密钥 进行补码（这个方案必须和js的方法是一样的）

@Editor robotyang at 2023
*/
func secretFill(secret []byte, pad byte, length int) []byte {
	if len(secret) >= length {
		return secret[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(secret))
	return append(pads, secret...)
}

/*
pKCS7UnFill 对明文 去除补码

@Editor robotyang at 2023
*/
func pKCS7UnFill(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}
