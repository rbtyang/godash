package dashtime

import (
	"time"
)

/*
@Editor robotyang at 2023

Cost is a ...
*/
func Cost(f func()) time.Duration {
	t1 := time.Now()
	f()
	return time.Since(t1)
}
