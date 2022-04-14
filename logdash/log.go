package logdash

import (
	"log"
)

func Debug(args ...interface{}) {
	log.Printf("[logdash.Debug] %#v", args)
}

func Debugf(format string, args ...interface{}) {
	log.Printf("[logdash.Debugf] "+format, args...)
}

func Info(args ...interface{}) {
	log.Printf("[logdash.Info] %#v", args)
}

func Infof(format string, args ...interface{}) {
	log.Printf("[logdash.Infof] "+format, args...)
}

func Warn(args ...interface{}) {
	log.Printf("[logdash.Warn] %#v", args)
}

func Warnf(format string, args ...interface{}) {
	log.Printf("[logdash.Warnf] "+format, args...)
}

func Error(args ...interface{}) {
	log.Printf("[logdash.Error] %#v", args)
}

func Errorf(format string, args ...interface{}) {
	log.Printf("[logdash.Errorf] "+format, args...)
}

func Panic(args ...interface{}) {
	log.Panicf("[logdash.Panic] %#v", args)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf("[logdash.Panicf] "+format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatalf("[logdash.Fatal] %#v", args)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf("[logdash.Fatalf] "+format, args...)
}
