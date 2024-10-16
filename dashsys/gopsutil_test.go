package dashsys_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/rbtyang/godash/dashsys"
)

/*
init is a ...
*/
func init() {
	log.Println("Before this tests")
}

/*
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
