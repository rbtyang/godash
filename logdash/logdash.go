package logdash

import (
	"context"
	"log"
)

func Debug(ctx context.Context, args ...interface{}) {
	log.Printf("[logdash.Debug] %#v", args)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	log.Printf("[logdash.Debugf] "+format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	log.Printf("[logdash.Info] %#v", args)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	log.Printf("[logdash.Infof] "+format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	log.Printf("[logdash.Warn] %#v", args)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	log.Printf("[logdash.Warnf] "+format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	log.Printf("[logdash.Error] %#v", args)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf("[logdash.Errorf] "+format, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	log.Panicf("[logdash.Panic] %#v", args)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	log.Panicf("[logdash.Panicf] "+format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	log.Fatalf("[logdash.Fatal] %#v", args)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	log.Fatalf("[logdash.Fatalf] "+format, args...)
}
