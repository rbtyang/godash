package dashlog_test

import (
	"github.com/rbtyang/godash/dashlog"
	"log"
	"testing"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
TestLog is a ...
*/
func TestLog(t *testing.T) {
	clsPre, _ := dashlog.Pre("哈哈哈")
	dashlog.Debug("1111", "aaaa", "bbbb")
	dashlog.Debug("1111", "aaaa", "bbbb")
	dashlog.Debugf("%v # %v # %v", "1111", "aaaa", "bbbb")
	dashlog.Info("1111", "aaaa", "bbbb")
	dashlog.Info("1111", "aaaa", "bbbb")
	clsPre() //or: defer clsPre()
	dashlog.Infof("%v # %v # %v", "1111", "aaaa", "bbbb")
	t.Logf("done: %v", true)
}
