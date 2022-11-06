package logdash_test

import (
	"github.com/rbtyang/godash/logdash"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestLog(t *testing.T) {
	clsPre, _ := logdash.Pre("哈哈哈")
	logdash.Debug("1111", "aaaa", "bbbb")
	logdash.Debug("1111", "aaaa", "bbbb")
	logdash.Debugf("%v # %v # %v", "1111", "aaaa", "bbbb")
	logdash.Info("1111", "aaaa", "bbbb")
	logdash.Info("1111", "aaaa", "bbbb")
	clsPre() //or: defer clsPre()
	logdash.Infof("%v # %v # %v", "1111", "aaaa", "bbbb")
	t.Logf("done: %v", true)
}
