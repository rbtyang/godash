package cryptdash

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func AesEncrypt(plaintext, secret []byte) (string, error) {
	ciphertext, err := AesCbcEncrypt(plaintext, secret)
	if err != nil {
		return "", err
	}
	cipherstr := base64.StdEncoding.EncodeToString(ciphertext)
	return cipherstr, nil
}

func AesDecrypt(cipherstr string, secret []byte) ([]byte, error) {
	ciphertext2, err := base64.StdEncoding.DecodeString(cipherstr)
	if err != nil {
		return nil, err
	}
	plaintext, err := AesCbcDecrypt(ciphertext2, secret)
	return plaintext, nil
}

// @param.secret key参数应该是AES密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256。
func AesCbcEncrypt(plaintext, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = pkcsFill(plaintext, blockSize)
	mode := cipher.NewCBCEncrypter(block, secret[:blockSize])
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func pkcsFill(plaintext []byte, blockSize int) []byte {
	fillNum := blockSize - len(plaintext)%blockSize
	fillText := bytes.Repeat([]byte{byte(fillNum)}, fillNum)
	return append(plaintext, fillText...)
}

func AesCbcDecrypt(ciphertext, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	mode := cipher.NewCBCDecrypter(block, secret[:blockSize])
	oritext := make([]byte, len(ciphertext))
	mode.CryptBlocks(oritext, ciphertext)
	oritext = pkcsUnFill(oritext)
	return oritext, nil
}

func pkcsUnFill(origData []byte) []byte {
	length := len(origData)
	unFillNum := int(origData[length-1])
	return origData[:(length - unFillNum)]
}

// ----------------------------

func AesEcbEncrypt(plaintext []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret[:aes.BlockSize])
	if err != nil {
		return nil, err
	}
	plaintext = zeroFill(plaintext, aes.BlockSize)

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(plaintext) > 0 {
		block.Encrypt(text, plaintext)
		plaintext = plaintext[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext, nil
}

func zeroFill(plaintext []byte, blockSize int) []byte {
	fillNum := blockSize - len(plaintext)%blockSize
	fillText := bytes.Repeat([]byte{0}, fillNum)
	return append(plaintext, fillText...)
}

func AesEcbDecrypt(ciphertext []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret[:aes.BlockSize])
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, 0)
	text := make([]byte, 16)
	for len(ciphertext) > 0 {
		block.Decrypt(text, ciphertext)
		ciphertext = ciphertext[aes.BlockSize:]
		plaintext = append(plaintext, text...)
	}
	plaintext = zeroUnFill(plaintext)

	return plaintext, nil
}

func zeroUnFill(plaintext []byte) []byte {
	plaintext = bytes.TrimFunc(plaintext, func(r rune) bool {
		return r == rune(0)
	})
	return plaintext
}

//---------------------------------------------------------------------------------------------------------------------

// AesDecryptJsHex
// ciphertext 密文（js的生成的密文后 进行了16进制的 hex.encoding，因此在调用该方法之前 go必须要进行 hex.DecodeString）
// secret 秘钥
// @reference https://mojotv.cn/go/crypto-js-with-golang
func AesDecryptJsHex(ciphertext, secret string) (string, error) {
	byt, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	str, err := AesDecryptJs(byt, []byte(secret))
	if err != nil {
		return "", err
	}
	return string(str), nil
}

// AesDecryptJs
// ciphertext 密文（js的生成的密文后 进行了16进制的 hex.encoding，因此在调用该方法之前 go必须要进行 hex.DecodeString）
// secret 秘钥
// @reference https://mojotv.cn/go/crypto-js-with-golang
func AesDecryptJs(ciphertext, secret []byte) ([]byte, error) {
	pkey := paddingLeft(secret, '0', 16) //和js的key补码方法一致

	block, err := aes.NewCipher(pkey) //选择加密算法
	if err != nil {
		return nil, fmt.Errorf("secret 长度必须 16/24/32长度: %s", err)
	}
	blockModel := cipher.NewCBCDecrypter(block, pkey) //和前端代码对应:   mode: CryptoJS.mode.CBC,// CBC算法
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = pKCS7UnPadding(plantText) //和前端代码对应:  padding: CryptoJS.pad.Pkcs7
	return plantText, nil
}

func pKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

//这个方案必须和js的方法是一样的
func paddingLeft(ori []byte, pad byte, length int) []byte {
	if len(ori) >= length {
		return ori[:length]
	}
	pads := bytes.Repeat([]byte{pad}, length-len(ori))
	return append(pads, ori...)
}

//---------------------------------------------------------------------------------------------------------------------
