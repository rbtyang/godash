package dashtime

import (
	"time"

	"github.com/rbtyang/godash/dashlog"
)

/*
Cost @Editor robotyang at 2023

# Cost 计算函数耗时

@Param fn 要执行的函数

@Return 函数耗时
*/
func Cost(fn func()) time.Duration {
	t1 := time.Now()
	fn()
	return time.Since(t1)
}

/*
CostPrint @Editor robotyang at 2023

# CostPrint 计算函数耗时 并打印到控制台

@Param title 打印时的函数说明标题

@Param fn 要执行的函数

@Param precision 计算精度，如 time.Second
*/
func CostPrint(title string, precision time.Duration, fn func()) {
	if title == "" {
		title = "Func"
	}
	var ts float64
	diff := Cost(fn)
	switch precision {
	case time.Hour:
		ts = diff.Hours()
		dashlog.Infof(title+" cost %v hour", ts)
	case time.Minute:
		ts = diff.Minutes()
		dashlog.Infof(title+" cost %v minute", ts)
	case time.Second:
		ts = diff.Seconds()
		dashlog.Infof(title+" cost %v second", ts)
	case time.Millisecond:
		ts = diff.Seconds()
		dashlog.Infof(title+" cost %v millisecond", ts)
	case time.Microsecond:
		ts = diff.Seconds()
		dashlog.Infof(title+" cost %v microsecond", ts)
	case time.Nanosecond:
		ts = diff.Seconds()
		dashlog.Infof(title+" cost %v nanosecond", ts)
	default:
		panic("dashtime.CostPrint: unsupported precision")
	}
}
