package dashlog

import (
	"fmt"
	"log"
)

// 用户 自定义前缀
var cuspre string

// 设置 自定义msg前缀

/*
Pre  @return.error always nil

@Editor robotyang at 2023
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

@Editor robotyang at 2023
*/
func Debug(args ...interface{}) {
	log.Printf(withPre("logdash.Debug")+"%#v", args)
}

/*
Debugf is a ...

@Editor robotyang at 2023
*/
func Debugf(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Debug")+format, args...)
}

/*
Info is a ...

@Editor robotyang at 2023
*/
func Info(args ...interface{}) {
	log.Printf(withPre("logdash.Info")+"%#v", args)
}

/*
Infof is a ...

@Editor robotyang at 2023
*/
func Infof(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Info")+format, args...)
}

/*
Warn is a ...

@Editor robotyang at 2023
*/
func Warn(args ...interface{}) {
	log.Printf(withPre("logdash.Warn")+"%#v", args)
}

/*
Warnf is a ...

@Editor robotyang at 2023
*/
func Warnf(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Warn")+format, args...)
}

/*
Error is a ...

@Editor robotyang at 2023
*/
func Error(args ...interface{}) {
	log.Printf(withPre("logdash.Error")+"%#v", args)
}

/*
Errorf is a ...

@Editor robotyang at 2023
*/
func Errorf(format string, args ...interface{}) {
	log.Printf(withPre("logdash.Error")+format, args...)
}

/*
Panic is a ...

@Editor robotyang at 2023
*/
func Panic(args ...interface{}) {
	log.Panicf(withPre("logdash.Panic")+"%#v", args)
}

/*
Panicf is a ...

@Editor robotyang at 2023
*/
func Panicf(format string, args ...interface{}) {
	log.Panicf(withPre("logdash.Panic")+format, args...)
}

/*
Fatal is a ...

@Editor robotyang at 2023
*/
func Fatal(args ...interface{}) {
	log.Fatalf(withPre("logdash.Fatal")+"%#v", args)
}

/*
Fatalf is a ...

@Editor robotyang at 2023
*/
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(withPre("logdash.Fatal")+format, args...)
}

/*
withPre is a ...

@Editor robotyang at 2023
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
