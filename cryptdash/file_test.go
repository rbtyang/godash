package cryptdash_test

import (
	"github.com/rbtyang/godash/cryptdash"
	"github.com/rbtyang/godash/randdash"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestFileEncrypt(t *testing.T) {
	{
		secret := []byte(randdash.Str(randdash.ModeNumAlphaSp, 32))

		err := cryptdash.FileEncryptByZyx("temp/test.txt", "temp/test_crypt.txt", secret)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("File Encrypt Success")

		err2 := cryptdash.FileDecryptByZyx("temp/test_crypt.txt", "temp/test_plain.txt", secret)
		if err2 != nil {
			t.Error(err2)
		}
		t.Log("File Decrypt Success")
	}
}