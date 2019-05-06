package log

import (
	"log"
	"os"
)

// Logger is an interface that the standard log package implements
type Logger interface {
	Printf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}

var DefaultLogger Logger

func mustGetDefaultLogger() Logger {
	if DefaultLogger == nil {
		DefaultLogger = NewDefaultLogger()
	}

	return DefaultLogger
}

func NewDefaultLogger() Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
}

func Printf(format string, v ...interface{}) {
	mustGetDefaultLogger().Printf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	mustGetDefaultLogger().Panicf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	mustGetDefaultLogger().Panicf(format, v...)
}
