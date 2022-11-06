package encodash_test

import (
	"github.com/rbtyang/godash/encodash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestBase64Encrypt(t *testing.T) {
	{
		data := "ZhangSan Ni Hao 123 哈哈"

		cipherstr := encodash.Base64Encrypt(data)
		t.Logf("cipherstr: %s", cipherstr)

		plainstr, err := encodash.Base64Decrypt(cipherstr)
		if err != nil {
			t.Error(err)
		}
		t.Logf("plainstr: %s", plainstr)

		assert.Equal(t, data, plainstr)
	}
}

func TestBase64UrlEncrypt(t *testing.T) {
	{
		data := "hello world12345!?$*&()'-@~"

		cipherstr0 := encodash.Base64Encrypt(data)
		t.Logf("cipherstr0: %s", cipherstr0)

		cipherstr1 := encodash.Base64UrlEncrypt(data)
		t.Logf("cipherstr1: %s", cipherstr1)

		plainstr, err := encodash.Base64UrlDecrypt(cipherstr1)
		if err != nil {
			t.Error(err)
		}
		t.Logf("plainstr: %s", plainstr)

		assert.Equal(t, data, plainstr)
	}
}
