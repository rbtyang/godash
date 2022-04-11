package logdash_test

import (
	"github.com/rbtyang/godash/logdash"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestTemp(t *testing.T) {
	logdash.Debug("1111", "aaaa", "bbbb")
	logdash.Debug("1111", "aaaa", "bbbb")
	logdash.Debugf("%v # %v # %v", "1111", "aaaa", "bbbb")
	logdash.Info("1111", "aaaa", "bbbb")
	logdash.Info("1111", "aaaa", "bbbb")
	logdash.Infof("%v # %v # %v", "1111", "aaaa", "bbbb")
	t.Logf("done: %v", true)
}
