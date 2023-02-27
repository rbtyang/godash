package sysinfodash

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

/*
CPUPercent is a ...

@Editor robotyang at 2023
*/
func CPUPercent() float64 {
	cc, err := cpu.Percent(time.Second, false)
	if len(cc) < 1 || err != nil {
		return float64(0)
	}
	return cc[0]
}

/*
MemoryPercent is a ...

@Editor robotyang at 2023
*/
func MemoryPercent() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return float64(0)
	}
	return v.UsedPercent
}

/*
SwapPercent is a ...

@Editor robotyang at 2023
*/
func SwapPercent() float64 {
	s, err := mem.SwapMemory()
	if err != nil {
		return float64(0)
	}
	return s.UsedPercent
}

/*
DiskPercent is a ...

@Editor robotyang at 2023
*/
func DiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
