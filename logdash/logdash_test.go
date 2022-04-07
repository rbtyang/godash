package logdash_test

import (
	"context"
	"github.com/rbtyang/godash/logdash"
	"log"
	"testing"
)

func init() {
	log.Println("Before this tests")
}

func TestTemp(t *testing.T) {
	ctx := context.Background()
	logdash.Debug(ctx, "1111", "aaaa", "bbbb")
	logdash.Debug(ctx, "1111", "aaaa", "bbbb")
	logdash.Debugf(ctx, "%v # %v # %v", "1111", "aaaa", "bbbb")
	logdash.Info(ctx, "1111", "aaaa", "bbbb")
	logdash.Info(ctx, "1111", "aaaa", "bbbb")
	logdash.Infof(ctx, "%v # %v # %v", "1111", "aaaa", "bbbb")
	t.Logf("done: %v", true)
}
