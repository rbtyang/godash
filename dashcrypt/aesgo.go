package dashcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

/*
AesBs64Encrypt @Editor robotyang at 2023

# AesBs64Encrypt 带Base64编码的 CBC模式加密（先走底层方法 AesCbcEncrypt 加密，再base64编码）；

@Param plaintext：明文；

@Param secret：密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256
*/
func AesBs64Encrypt(plaintext, secret []byte) (string, error) {
	ciphertext, err := AesCbcEncrypt(plaintext, secret)
	if err != nil {
		return "", err
	}
	cipherstr := base64.StdEncoding.EncodeToString(ciphertext)
	return cipherstr, nil
}

/*
AesBs64Decrypt @Editor robotyang at 2023

# AesBs64Decrypt 带Base64编码的 CBC模式解密（先base64解码，再走底层方法 AesCbcEncrypt 解密）；

@Param cipherstr：明文；

@Param secret：密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256
*/
func AesBs64Decrypt(cipherstr string, secret []byte) ([]byte, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cipherstr)
	if err != nil {
		return nil, err
	}
	plaintext, err := AesCbcDecrypt(ciphertext, secret)
	return plaintext, nil
}

//-------------------------------------------------------------------------------------------------------

/*
AesCbcEncrypt @Editor robotyang at 2023

# AesCbcEncrypt CBC模式加密；

@Param plaintext：明文；

@Param secret：密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256；
*/
func AesCbcEncrypt(plaintext, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = aesPkcsFill(plaintext, blockSize)
	mode := cipher.NewCBCEncrypter(block, secret[:blockSize])
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

/*
AesCbcDecrypt @Editor robotyang at 2023

# AesCbcDecrypt CBC模式解密；

@Param ciphertext：密文；

@Param secret：密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256；
*/
func AesCbcDecrypt(ciphertext, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	mode := cipher.NewCBCDecrypter(block, secret[:blockSize])
	oritext := make([]byte, len(ciphertext))
	mode.CryptBlocks(oritext, ciphertext)
	oritext = aesPkcsUnFill(oritext)
	return oritext, nil
}

/*
aesZeroFill @Editor robotyang at 2023

# aesZeroFill 填充补码
*/
func aesPkcsFill(plaintext []byte, blockSize int) []byte {
	fillNum := blockSize - len(plaintext)%blockSize
	fillText := bytes.Repeat([]byte{byte(fillNum)}, fillNum)
	return append(plaintext, fillText...)
}

/*
aesPkcsUnFill @Editor robotyang at 2023

# aesPkcsUnFill 去除补码
*/
func aesPkcsUnFill(origData []byte) []byte {
	length := len(origData)
	unFillNum := int(origData[length-1])
	return origData[:(length - unFillNum)]
}

//-------------------------------------------------------------------------------------------------------

/*
AesEcbEncrypt @Editor robotyang at 2023

# AesEcbEncrypt ECB模式加密；

@Param plaintext：明文

@Param secret：密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256；
*/
func AesEcbEncrypt(plaintext []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret[:aes.BlockSize])
	if err != nil {
		return nil, err
	}
	plaintext = aesZeroFill(plaintext, aes.BlockSize)

	ciphertext := make([]byte, 0)
	text := make([]byte, 16)
	for len(plaintext) > 0 {
		block.Encrypt(text, plaintext)
		plaintext = plaintext[aes.BlockSize:]
		ciphertext = append(ciphertext, text...)
	}
	return ciphertext, nil
}

/*
AesEcbDecrypt @Editor robotyang at 2023

# AesEcbDecrypt ECB模式解密；

@Param ciphertext：密文；

@Param secret：密钥，可以是16、24或32字节，用以选择AES-128、AES-192或AES-256；
*/
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
	plaintext = aesZeroUnFill(plaintext)

	return plaintext, nil
}

/*
aesZeroFill @Editor robotyang at 2023

# aesZeroFill 填充补码
*/
func aesZeroFill(plaintext []byte, blockSize int) []byte {
	fillNum := blockSize - len(plaintext)%blockSize
	fillText := bytes.Repeat([]byte{0}, fillNum)
	return append(plaintext, fillText...)
}

/*
aesZeroUnFill @Editor robotyang at 2023

# aesZeroUnFill 去除补码
*/
func aesZeroUnFill(plaintext []byte) []byte {
	plaintext = bytes.TrimFunc(plaintext, func(r rune) bool {
		return r == rune(0)
	})
	return plaintext
}
