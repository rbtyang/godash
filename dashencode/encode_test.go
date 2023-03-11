package dashencode_test

import (
	"github.com/rbtyang/godash/dashencode"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before this tests")
}

/*
TestBase64Encrypt is a ...

@Editor robotyang at 2023
*/
func TestBase64Encrypt(t *testing.T) {
	{
		data := "ZhangSan Ni Hao 123 哈哈"

		cipherstr := dashencode.Base64Encode(data)
		t.Logf("cipherstr: %s", cipherstr)

		plainstr, err := dashencode.Base64Decode(cipherstr)
		if err != nil {
			t.Error(err)
		}
		t.Logf("plainstr: %s", plainstr)

		assert.Equal(t, data, plainstr)
	}
}

/*
TestBase64UrlEncrypt is a ...

@Editor robotyang at 2023
*/
func TestBase64UrlEncrypt(t *testing.T) {
	{
		data := "hello world12345!?$*&()'-@~"

		cipherstr0 := dashencode.Base64Encode(data)
		t.Logf("cipherstr0: %s", cipherstr0)

		cipherstr1 := dashencode.Base64UrlEncode(data)
		t.Logf("cipherstr1: %s", cipherstr1)

		plainstr, err := dashencode.Base64UrlDecode(cipherstr1)
		if err != nil {
			t.Error(err)
		}
		t.Logf("plainstr: %s", plainstr)

		assert.Equal(t, data, plainstr)
	}
}
