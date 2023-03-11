package dashsys_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashsys"
	"log"
	"testing"
	"time"
)

/*
init is a ...

@Editor robotyang at 2023
*/
func init() {
	log.Println("Before this tests")
}

/*
TestSysinfoPercent is a ...

@Editor robotyang at 2023
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
