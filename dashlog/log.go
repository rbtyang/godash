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
Debug is a ...
*/
func Debug(args ...interface{}) {
	prefix := withPre("dashlog.Debug", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Debugf is a ...
*/
func Debugf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Debug", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Info is a ...
*/
func Info(args ...interface{}) {
	prefix := withPre("dashlog.Info", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Infof is a ...
*/
func Infof(format string, args ...interface{}) {
	prefix := withPre("dashlog.Info", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Warn is a ...
*/
func Warn(args ...interface{}) {
	prefix := withPre("dashlog.Warn", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Warnf is a ...
*/
func Warnf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Warn", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Error is a ...
*/
func Error(args ...interface{}) {
	prefix := withPre("dashlog.Error", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Errorf is a ...
*/
func Errorf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Error", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Panic is a ...
*/
func Panic(args ...interface{}) {
	prefix := withPre("dashlog.Panic", dashrun.LastCallerFuncName(true))
	log.Panicf(prefix+"%#v", args)
}

/*
Panicf is a ...
*/
func Panicf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Panic", dashrun.LastCallerFuncName(true))
	log.Panicf(prefix+format, args...)
}

/*
Fatal is a ...
*/
func Fatal(args ...interface{}) {
	prefix := withPre("dashlog.Fatal", dashrun.LastCallerFuncName(true))
	log.Fatalf(prefix+"%#v", args)
}

/*
Fatalf is a ...
*/
func Fatalf(format string, args ...interface{}) {
	prefix := withPre("dashlog.Fatal", dashrun.LastCallerFuncName(true))
	log.Fatalf(prefix+format, args...)
}

/*
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
