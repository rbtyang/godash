package dashsys_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashsys"
	"log"
	"testing"
	"time"
)

/*
@Editor robotyang at 2023

init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
@Editor robotyang at 2023

TestSysinfoPercent is a ...
*/
func TestSysinfoPercent(t *testing.T) {
	loop := 10
	for loop > 0 {
		loop--
		cpu := dashsys.CPUPercent()
		mem := dashsys.MemoryPercent()
		swa := dashsys.SwapPercent()
		disk := dashsys.DiskPercent()
		fmt.Printf("CPU:%8.4f,  Mem:%8.4f,  Swa:%8.4f,  Disk:%8.4f\n", cpu, mem, swa, disk)
		time.Sleep(time.Millisecond)
	}
}
