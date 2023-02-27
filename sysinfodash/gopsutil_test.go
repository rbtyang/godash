package sysinfodash_test

import (
	"fmt"
	"github.com/rbtyang/godash/sysinfodash"
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
		cpu := sysinfodash.CPUPercent()
		mem := sysinfodash.MemoryPercent()
		swa := sysinfodash.SwapPercent()
		disk := sysinfodash.DiskPercent()
		fmt.Printf("CPU:%8.4f,  Mem:%8.4f,  Swa:%8.4f,  Disk:%8.4f\n", cpu, mem, swa, disk)
		time.Sleep(time.Millisecond)
	}
}
