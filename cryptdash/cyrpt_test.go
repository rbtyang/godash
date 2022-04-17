package cryptdash_test

import (
	"github.com/rbtyang/godash/cryptdash"
	"github.com/rbtyang/godash/randdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestAesEncrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := []byte(randdash.Str(randdash.ModeNumAlphaSp, 32))

		ciphertext, err :=  cryptdash.AesEncrypt(data, secret)
		if err != nil {
			t.Error(err)
			return
		}
		plaintext, err := cryptdash.AesDecrypt(ciphertext, secret)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, data, plaintext)
	}
}

func TestAesCbcEncrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := randdash.Str(randdash.ModeNumAlphaSp, 32)

		ciphertext, err :=  cryptdash.AesCbcEncrypt(data, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}
		plaintext, err := cryptdash.AesCbcDecrypt(ciphertext, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, data, plaintext)
	}
}

func TestAesEcbEncrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := "7f797e9a4e2c3a9b190225d299214ce4"

		ciphertext, err :=  cryptdash.AesEcbEncrypt(data, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}
		plaintext, err := cryptdash.AesEcbDecrypt(ciphertext, []byte(secret))
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, data, plaintext)
	}
}