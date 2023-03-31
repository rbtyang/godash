package dashlog

import (
	"fmt"
	"github.com/rbtyang/godash/dashrun"
	"log"
)

// 用户 自定义前缀
var cuspre string

// 设置 自定义msg前缀

/*
@Editor robotyang at 2023

Pre 设置日志前缀

@Param pre 日志前缀字符串

@Return.error always nil
*/
func Pre(pre string) (func(), error) {
	cuspre = pre
	return func() {
		// 去除 自定义msg前缀
		cuspre = ""
	}, nil
}

/*
@Editor robotyang at 2023

Debug is a ...
*/
func Debug(args ...interface{}) {
	prefix := withPre("dashlog.Debug", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
@Editor robotyang at 2023

Debugf is a ...
*/
func Debugf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Debug", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
@Editor robotyang at 2023

Info is a ...
*/
func Info(args ...interface{}) {
	prefix := withPre("dashlog.Info", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
@Editor robotyang at 2023

Infof is a ...
*/
func Infof(format string, args ...interface{}) {
	prefix := withPre("dashlog.Info", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
@Editor robotyang at 2023

Warn is a ...
*/
func Warn(args ...interface{}) {
	prefix := withPre("dashlog.Warn", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
@Editor robotyang at 2023

Warnf is a ...
*/
func Warnf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Warn", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
@Editor robotyang at 2023

Error is a ...
*/
func Error(args ...interface{}) {
	prefix := withPre("dashlog.Error", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
@Editor robotyang at 2023

Errorf is a ...
*/
func Errorf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Error", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
@Editor robotyang at 2023

Panic is a ...
*/
func Panic(args ...interface{}) {
	prefix := withPre("dashlog.Panic", dashrun.LastCallerFuncName(true))
	log.Panicf(prefix+"%#v", args)
}

/*
@Editor robotyang at 2023

Panicf is a ...
*/
func Panicf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Panic", dashrun.LastCallerFuncName(true))
	log.Panicf(prefix+format, args...)
}

/*
@Editor robotyang at 2023

Fatal is a ...
*/
func Fatal(args ...interface{}) {
	prefix := withPre("dashlog.Fatal", dashrun.LastCallerFuncName(true))
	log.Fatalf(prefix+"%#v", args)
}

/*
@Editor robotyang at 2023

Fatalf is a ...
*/
func Fatalf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Fatal", dashrun.LastCallerFuncName(true))
	log.Fatalf(prefix+format, args...)
}

/*
@Editor robotyang at 2023

withPre is a ...
*/
func withPre(levelPre, lastFunc string) string {
	var prefix string
	if cuspre == "" {
		prefix = fmt.Sprintf("[%v] [%v] ", levelPre, lastFunc)
	} else {
		prefix = fmt.Sprintf("[%v] [%v] [%v] ", levelPre, cuspre, lastFunc)
	}
	return prefix
}
