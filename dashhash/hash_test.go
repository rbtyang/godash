package dashhash_test

import (
	"github.com/rbtyang/godash/dashhash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	{
		want := "7f797e9a4e2c3a9b190225d299214ce4"
		enStr := dashhash.Md5("hello world 123")
		assert.Equal(t, want, enStr)
	}
	{
		want := "7f797e9a4e2c3a9b190225d299214ce4"
		enStr := dashhash.Md5ByteToStr([]byte("hello world 123"))
		assert.Equal(t, want, enStr)
	}
	{
		want := "078175f6152e6e900c28facbbdfe8b1911db8f1e"
		enStr := dashhash.Sha1("hello world 123")
		assert.Equal(t, want, enStr)
	}
	{
		want := "d4223bf93e202505a6a501421a88d9fa43341f7757e217dd603ccdce157c13bd"
		enStr := dashhash.Sha256("hello world 123")
		assert.Equal(t, want, enStr)
	}
	{
		want := "32a47d7816de56bc617d1a87425ae77fc3556dc96f79e88884601c7100224a35dfe63fb53b75e87d0fdfbc0a7b8c6d9fca3c78d73e448bee759610ac5450edd6"
		enStr := dashhash.Sha512("hello world 123")
		assert.Equal(t, want, enStr)
	}
	{
		want := "f1b90b4efd0e5c7db52dfa0efd6521a3"
		enStr := dashhash.Hmac("key2", "hello")
		assert.Equal(t, want, enStr)
	}
}
