package log

import (
	"fmt"
	"github.com/iesreza/gutil/logger"
	"os"
)

var log = logger.New()

// Fatal is just like func l.Critical logger except that it is followed by exit to program
func Fatal(format interface{}, a ...interface{}) {

	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
	os.Exit(1)
}

// FatalF is just like func l.CriticalF logger except that it is followed by exit to program
func FatalF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
	os.Exit(1)
}

// FatalF is just like func l.CriticalF logger except that it is followed by exit to program
func Fatalf(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
	os.Exit(1)
}

// Panic is just like func l.Critical except that it is followed by a call to panic
func Panic(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
	panic(fmt.Sprintf(fmt.Sprint(format), a...))
}

// PanicF is just like func l.CriticalF except that it is followed by a call to panic
func PanicF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
	panic(fmt.Sprintf(fmt.Sprint(format), a...))
}

// PanicF is just like func l.CriticalF except that it is followed by a call to panic
func Panicf(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
	panic(fmt.Sprintf(fmt.Sprint(format), a...))
}

// Critical logs a message at a Critical Level
func Critical(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// CriticalF logs a message at Critical level using the same syntax and options as fmt.Printf
func CriticalF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// CriticalF logs a message at Critical level using the same syntax and options as fmt.Printf
func Criticalf(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.CriticalLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// Error logs a message at Error level
func Error(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.ErrorLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// ErrorF logs a message at Error level using the same syntax and options as fmt.Printf
func ErrorF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.ErrorLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// ErrorF logs a message at Error level using the same syntax and options as fmt.Printf
func Errorf(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.ErrorLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// Warning logs a message at Warning level
func Warning(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.WarningLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// WarningF logs a message at Warning level using the same syntax and options as fmt.Printf
func WarningF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.WarningLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// WarningF logs a message at Warning level using the same syntax and options as fmt.Printf
func Warningf(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.WarningLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// Notice logs a message at Notice level
func Notice(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.NoticeLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// NoticeF logs a message at Notice level using the same syntax and options as fmt.Printf
func NoticeF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.NoticeLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// NoticeF logs a message at Notice level using the same syntax and options as fmt.Printf
func Noticef(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.NoticeLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// Info logs a message at Info level
func Info(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.InfoLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// InfoF logs a message at Info level using the same syntax and options as fmt.Printf
func InfoF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.InfoLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// InfoF logs a message at Info level using the same syntax and options as fmt.Printf
func Infof(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.InfoLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// Debug logs a message at Debug level
func Debug(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.DebugLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// DebugF logs a message at Debug level using the same syntax and options as fmt.Printf
func DebugF(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.DebugLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}

// DebugF logs a message at Debug level using the same syntax and options as fmt.Printf
func Debugf(format interface{}, a ...interface{}) {
	log.LogWrapped(logger.DebugLevel, fmt.Sprintf(fmt.Sprint(format), a...))
}
