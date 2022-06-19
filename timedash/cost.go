package timedash

import (
	"time"
)

func Cost(f func()) time.Duration {
	t1 := time.Now()
	f()
	return time.Since(t1)
}
