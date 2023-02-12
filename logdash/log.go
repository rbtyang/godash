package logdash

import (
	"fmt"
	"log"
)

// 用户 自定义前缀
var cuspre string

// 设置 自定义msg前缀
// @return.error always nil
func Pre(pre string) (func(), error) {
	cuspre = pre
	return func() {
		// 去除 自定义msg前缀
		cuspre = ""
	}, nil
}

func Debug(args ...interface{}) {
	log.Printf(withPre("logdash.Debug")+"%#v", args)
}

func Debugf(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Debug")+format, args...)
}

func Info(args ...interface{}) {
	log.Printf(withPre("logdash.Info")+"%#v", args)
}

func Infof(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Info")+format, args...)
}

func Warn(args ...interface{}) {
	log.Printf(withPre("logdash.Warn")+"%#v", args)
}

func Warnf(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Warn")+format, args...)
}

func Error(args ...interface{}) {
	log.Printf(withPre("logdash.Error")+"%#v", args)
}

func Errorf(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Error")+format, args...)
}

func Panic(args ...interface{}) {
	log.Panicf(withPre("logdash.Panic")+"%#v", args)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(withPre("logdash.Panic")+format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatalf(withPre("logdash.Fatal")+"%#v", args)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(withPre("logdash.Fatal")+format, args...)
}

func withPre(dashpre string) string {
	var prefix string
	if cuspre == "" {
		prefix = fmt.Sprintf("[%v] ", dashpre)
	} else {
		prefix = fmt.Sprintf("[%v] [%v] ", dashpre, cuspre)
	}
	return prefix
}
