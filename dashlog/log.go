package dashlog

import (
	"fmt"
	"log"
)

// 用户 自定义前缀
var cuspre string

// 设置 自定义msg前缀

/*
@Editor robotyang at 2023

Pre  @return.error always nil
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
	log.Printf(withPre("dashlog.Debug")+"%#v", args)
}

/*
@Editor robotyang at 2023

Debugf is a ...
*/
func Debugf(format string, args ...interface{}) {
	log.Printf(withPre("dashlog.Debug")+format, args...)
}

/*
@Editor robotyang at 2023

Info is a ...
*/
func Info(args ...interface{}) {
	log.Printf(withPre("dashlog.Info")+"%#v", args)
}

/*
@Editor robotyang at 2023

Infof is a ...
*/
func Infof(format string, args ...interface{}) {
	log.Printf(withPre("dashlog.Info")+format, args...)
}

/*
@Editor robotyang at 2023

Warn is a ...
*/
func Warn(args ...interface{}) {
	log.Printf(withPre("dashlog.Warn")+"%#v", args)
}

/*
@Editor robotyang at 2023

Warnf is a ...
*/
func Warnf(format string, args ...interface{}) {
	log.Printf(withPre("dashlog.Warn")+format, args...)
}

/*
@Editor robotyang at 2023

Error is a ...
*/
func Error(args ...interface{}) {
	log.Printf(withPre("dashlog.Error")+"%#v", args)
}

/*
@Editor robotyang at 2023

Errorf is a ...
*/
func Errorf(format string, args ...interface{}) {
	log.Printf(withPre("dashlog.Error")+format, args...)
}

/*
@Editor robotyang at 2023

Panic is a ...
*/
func Panic(args ...interface{}) {
	log.Panicf(withPre("dashlog.Panic")+"%#v", args)
}

/*
@Editor robotyang at 2023

Panicf is a ...
*/
func Panicf(format string, args ...interface{}) {
	log.Panicf(withPre("dashlog.Panic")+format, args...)
}

/*
@Editor robotyang at 2023

Fatal is a ...
*/
func Fatal(args ...interface{}) {
	log.Fatalf(withPre("dashlog.Fatal")+"%#v", args)
}

/*
@Editor robotyang at 2023

Fatalf is a ...
*/
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(withPre("dashlog.Fatal")+format, args...)
}

/*
@Editor robotyang at 2023

withPre is a ...
*/
func withPre(dashpre string) string {
	var prefix string
	if cuspre == "" {
		prefix = fmt.Sprintf("[%v] ", dashpre)
	} else {
		prefix = fmt.Sprintf("[%v] [%v] ", dashpre, cuspre)
	}
	return prefix
}
