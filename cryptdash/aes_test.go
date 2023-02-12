package cryptdash_test

import (
	"encoding/hex"
	"github.com/rbtyang/godash/cryptdash"
	"github.com/rbtyang/godash/randdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}
func TestAesDecryptJs(t *testing.T) {
	wanttext := "123123阿斯蒂芬!@#asdasd"

	{ //mixed plaintext 1
		secret := "ABCDEF1234123412"
		ciphertext := "9e2819811c90f03c407ecfc7253b240556ff169e3371127c46a651ae1920d8df"
		bs, err := hex.DecodeString(ciphertext)
		if err != nil {
			t.Error(err)
		}
		recvbyt, err := cryptdash.AesDecryptJs(bs, []byte(secret))
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, string(recvbyt), wanttext)
	}
	{ //mixed plaintext 1
		secret := "ABCDEF1234123412"
		ciphertext := "9e2819811c90f03c407ecfc7253b240556ff169e3371127c46a651ae1920d8df"
		recvtext, err := cryptdash.AesDecryptJsHex(ciphertext, secret)
		if err != nil {
			return
		}
		assert.Equal(t, recvtext, wanttext)
	}
	{ //mixed plaintext 2 diff secret
		secret := "M8xMxeX6rgBsveTF"
		ciphertext := "3c258944eb76163936f7b01aa453425f28110783c85d6241b319aa24c25e9023"
		recvtext, err := cryptdash.AesDecryptJsHex(ciphertext, secret)
		if err != nil {
			return
		}
		assert.Equal(t, recvtext, wanttext)
	}
	{ //mixed plaintext 3 long secret
		secret := "M8xMxeX6rgBsveTFM8xMxeX6rgBsveTF"
		ciphertext := "3c258944eb76163936f7b01aa453425f28110783c85d6241b319aa24c25e9023"
		recvtext, err := cryptdash.AesDecryptJsHex(ciphertext, secret)
		if err != nil {
			return
		}
		assert.Equal(t, recvtext, wanttext)
	}
	
	t.Log("done")
}

func TestAesEncrypt(t *testing.T) {
	{
		data := []byte("Hello World 123 哈哈")
		secret := []byte(randdash.Str(randdash.ModeNumAlphaSp, 32))

		ciphertext, err := cryptdash.AesEncrypt(data, secret)
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

		ciphertext, err := cryptdash.AesCbcEncrypt(data, []byte(secret))
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

		ciphertext, err := cryptdash.AesEcbEncrypt(data, []byte(secret))
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
