package dashcrypt_test

import (
	"log"
	"testing"

	"github.com/rbtyang/godash/dashcrypt"
	"github.com/rbtyang/godash/dashrand"
	"github.com/stretchr/testify/assert"
)

/*
init is a ...
*/
func init() {
	log.Println("Before aesgo_test.go tests")
}

/*
TestAesBs64Encrypt is a ...
*/
func TestAesBs64Encrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := []byte(dashrand.Str(dashrand.ModeNumAlphaSp, 32))

		ciphertext, err := dashcrypt.AesBs64Encrypt(data, secret)
		if err != nil {
			t.Error(err)
			return
		}
		plaintext, err := dashcrypt.AesBs64Decrypt(ciphertext, secret)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, data, plaintext)
	}
}

/*
TestAesCbcEncrypt is a ...
*/
func TestAesCbcEncrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := dashrand.Str(dashrand.ModeNumAlphaSp, 32)

		ciphertext, err := dashcrypt.AesCbcEncrypt(data, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}
		plaintext, err := dashcrypt.AesCbcDecrypt(ciphertext, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, data, plaintext)
	}
}

/*
TestAesEcbEncrypt is a ...
*/
func TestAesEcbEncrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := "7f797e9a4e2c3a9b190225d299214ce4"

		ciphertext, err := dashcrypt.AesEcbEncrypt(data, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}
		plaintext, err := dashcrypt.AesEcbDecrypt(ciphertext, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, data, plaintext)
	}
}
