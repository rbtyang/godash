package dashsys

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

/*
@Editor robotyang at 2023

CPUPercent is a ...
*/
func CPUPercent() float64 {
	cc, err := cpu.Percent(time.Second, false)
	if len(cc) < 1 || err != nil {
		return float64(0)
	}
	return cc[0]
}

/*
@Editor robotyang at 2023

MemoryPercent is a ...
*/
func MemoryPercent() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		return float64(0)
	}
	return v.UsedPercent
}

/*
@Editor robotyang at 2023

SwapPercent is a ...
*/
func SwapPercent() float64 {
	s, err := mem.SwapMemory()
	if err != nil {
		return float64(0)
	}
	return s.UsedPercent
}

/*
@Editor robotyang at 2023

DiskPercent is a ...
*/
func DiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}
