package dashende_test

import (
	"log"
	"testing"

	"github.com/rbtyang/godash/dashende"
	"github.com/stretchr/testify/assert"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestBase64Encrypt is a ...
*/
func TestBase64Encrypt(t *testing.T) {
	{
		data := "ZhangSan Ni Hao 123 哈哈"

		cipherstr := dashende.Base64Encode(data)
		t.Logf("cipherstr: %s", cipherstr)

		plainstr, err := dashende.Base64Decode(cipherstr)
		if err != nil {
			t.Error(err)
		}
		t.Logf("plainstr: %s", plainstr)

		assert.Equal(t, data, plainstr)
	}
}

/*
TestBase64UrlEncrypt is a ...
*/
func TestBase64UrlEncrypt(t *testing.T) {
	{
		data := "hello world12345!?$*&()'-@~"

		cipherstr0 := dashende.Base64Encode(data)
		t.Logf("cipherstr0: %s", cipherstr0)

		cipherstr1 := dashende.Base64UrlEncode(data)
		t.Logf("cipherstr1: %s", cipherstr1)

		plainstr, err := dashende.Base64UrlDecode(cipherstr1)
		if err != nil {
			t.Error(err)
		}
		t.Logf("plainstr: %s", plainstr)

		assert.Equal(t, data, plainstr)
	}
}
