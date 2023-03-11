package dashtime

import (
	"time"
)

/*
Cost is a ...

@Editor robotyang at 2023
*/
func Cost(f func()) time.Duration {
	t1 := time.Now()
	f()
	return time.Since(t1)
}
