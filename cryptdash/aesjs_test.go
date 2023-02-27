package cryptdash_test

import (
	"encoding/hex"
	"github.com/rbtyang/godash/cryptdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before aesjs_test.go tests")
}

/*
TestAesDecryptJs is a ...

@Editor robotyang at 2023
*/
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
