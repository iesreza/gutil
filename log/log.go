package log

import (
	"fmt"
	"github.com/iesreza/gutil/logger"
	"os"
)

var log = logger.New()

// Fatal is just like func l.Critical logger except that it is followed by exit to program
func Fatal(message string) {
	log.Fatal(message)
	os.Exit(1)
}

// FatalF is just like func l.CriticalF logger except that it is followed by exit to program
func FatalF(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(format, a...))
	os.Exit(1)
}

// FatalF is just like func l.CriticalF logger except that it is followed by exit to program
func Fatalf(format string, a ...interface{}) {
	log.Fatal(fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Panic is just like func l.Critical except that it is followed by a call to panic
func Panic(message string) {
	log.Panic(message)
	panic(message)
}

// PanicF is just like func l.CriticalF except that it is followed by a call to panic
func PanicF(format string, a ...interface{}) {
	log.PanicF(fmt.Sprintf(format, a...))
	panic(fmt.Sprintf(format, a...))
}

// PanicF is just like func l.CriticalF except that it is followed by a call to panic
func Panicf(format string, a ...interface{}) {
	log.PanicF(fmt.Sprintf(format, a...))
	panic(fmt.Sprintf(format, a...))
}

// Critical logs a message at a Critical Level
func Critical(message string) {
	log.Critical(message)
}

// CriticalF logs a message at Critical level using the same syntax and options as fmt.Printf
func CriticalF(format string, a ...interface{}) {
	log.CriticalF(fmt.Sprintf(format, a...))
}

// CriticalF logs a message at Critical level using the same syntax and options as fmt.Printf
func Criticalf(format string, a ...interface{}) {
	log.CriticalF(fmt.Sprintf(format, a...))
}

// Error logs a message at Error level
func Error(message string) {
	log.Error(message)
}

// ErrorF logs a message at Error level using the same syntax and options as fmt.Printf
func ErrorF(format string, a ...interface{}) {
	log.ErrorF(fmt.Sprintf(format, a...))
}

// ErrorF logs a message at Error level using the same syntax and options as fmt.Printf
func Errorf(format string, a ...interface{}) {
	log.ErrorF(fmt.Sprintf(format, a...))
}

// Warning logs a message at Warning level
func Warning(message string) {
	log.Warning(message)
}

// WarningF logs a message at Warning level using the same syntax and options as fmt.Printf
func WarningF(format string, a ...interface{}) {
	log.Warningf(fmt.Sprintf(format, a...))
}

// WarningF logs a message at Warning level using the same syntax and options as fmt.Printf
func Warningf(format string, a ...interface{}) {
	log.Warningf(fmt.Sprintf(format, a...))
}

// Notice logs a message at Notice level
func Notice(message string) {
	log.Notice(message)
}

// NoticeF logs a message at Notice level using the same syntax and options as fmt.Printf
func NoticeF(format string, a ...interface{}) {
	log.NoticeF(fmt.Sprintf(format, a...))
}

// NoticeF logs a message at Notice level using the same syntax and options as fmt.Printf
func Noticef(format string, a ...interface{}) {
	log.Noticef(fmt.Sprintf(format, a...))
}

// Info logs a message at Info level
func Info(message string) {
	log.Info(message)
}

// InfoF logs a message at Info level using the same syntax and options as fmt.Printf
func InfoF(format string, a ...interface{}) {
	log.InfoF(fmt.Sprintf(format, a...))
}

// InfoF logs a message at Info level using the same syntax and options as fmt.Printf
func Infof(format string, a ...interface{}) {
	log.InfoF(fmt.Sprintf(format, a...))
}

// Debug logs a message at Debug level
func Debug(message string) {
	log.Debug(message)
}

// DebugF logs a message at Debug level using the same syntax and options as fmt.Printf
func DebugF(format string, a ...interface{}) {
	log.DebugF(fmt.Sprintf(format, a...))
}

// DebugF logs a message at Debug level using the same syntax and options as fmt.Printf
func Debugf(format string, a ...interface{}) {
	log.DebugF(fmt.Sprintf(format, a...))
}
