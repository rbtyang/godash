package dashlog

import (
	"fmt"
	"log"

	"github.com/rbtyang/godash/dashrun"
)

// 用户 自定义前缀
var cuspre string

/*
Pre @Editor robotyang at 2023

# Pre 设置日志前缀

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
Debug @Editor robotyang at 2023

# Debug is a ...
*/
func Debug(args ...any) {
	prefix := withPre("dashlog.Debug", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Debugf @Editor robotyang at 2023

# Debugf is a ...
*/
func Debugf(format string, args ...any) {
	prefix := withPre("dashlog.Debug", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Info @Editor robotyang at 2023

# Info is a ...
*/
func Info(args ...any) {
	prefix := withPre("dashlog.Info", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Infof @Editor robotyang at 2023

# Infof is a ...
*/
func Infof(format string, args ...any) {
	prefix := withPre("dashlog.Info", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Warn @Editor robotyang at 2023

# Warn is a ...
*/
func Warn(args ...any) {
	prefix := withPre("dashlog.Warn", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Warnf @Editor robotyang at 2023

# Warnf is a ...
*/
func Warnf(format string, args ...any) {
	prefix := withPre("dashlog.Warn", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Error @Editor robotyang at 2023

# Error is a ...
*/
func Error(args ...any) {
	prefix := withPre("dashlog.Error", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+"%#v", args)
}

/*
Errorf @Editor robotyang at 2023

# Errorf is a ...
*/
func Errorf(format string, args ...any) {
	prefix := withPre("dashlog.Error", dashrun.LastCallerFuncName(true))
	log.Printf(prefix+format, args...)
}

/*
Panic @Editor robotyang at 2023

# Panic is a ...
*/
func Panic(args ...any) {
	prefix := withPre("dashlog.Panic", dashrun.LastCallerFuncName(true))
	log.Panicf(prefix+"%#v", args)
}

/*
Panicf @Editor robotyang at 2023

# Panicf is a ...
*/
func Panicf(format string, args ...any) {
	prefix := withPre("dashlog.Panic", dashrun.LastCallerFuncName(true))
	log.Panicf(prefix+format, args...)
}

/*
Fatal @Editor robotyang at 2023

# Fatal is a ...
*/
func Fatal(args ...any) {
	prefix := withPre("dashlog.Fatal", dashrun.LastCallerFuncName(true))
	log.Fatalf(prefix+"%#v", args)
}

/*
Fatalf @Editor robotyang at 2023

# Fatalf is a ...
*/
func Fatalf(format string, args ...any) {
	prefix := withPre("dashlog.Fatal", dashrun.LastCallerFuncName(true))
	log.Fatalf(prefix+format, args...)
}

/*
withPre @Editor robotyang at 2023

# withPre is a ...
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
