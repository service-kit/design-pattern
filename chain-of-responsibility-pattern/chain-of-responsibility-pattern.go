package chain_of_responsibility_pattern

import "fmt"

const (
	INFO  = 1
	DEBUG = 2
	ERROR = 3
)

type IAbstractLogger interface {
	SetNextLogger(nl IAbstractLogger)
	LogMessage(level int, msg string)
	SetWriter(func(string))
}

type AbstractLogger struct {
	level      int
	nextLogger IAbstractLogger
	writer     func(string)
}

func (a *AbstractLogger) SetNextLogger(nl IAbstractLogger) {
	a.nextLogger = nl
}

func (a *AbstractLogger) LogMessage(level int, msg string) {
	if a.level <= level {
		a.writer(msg)
	}
	if nil != a.nextLogger {
		a.nextLogger.LogMessage(level, msg)
	}
}

func (a *AbstractLogger) SetWriter(f func(string)) {
	a.writer = f
}

func NewConsoleLogger() IAbstractLogger {
	logger := &AbstractLogger{level: INFO}
	logger.SetWriter(func(s string) {
		fmt.Println("Console Logger:", s)
	})
	return logger
}

func NewFileLogger() IAbstractLogger {
	logger := &AbstractLogger{level: DEBUG}
	logger.SetWriter(func(s string) {
		fmt.Println("File Logger:", s)
	})
	return logger
}

func NewErrorLogger() IAbstractLogger {
	logger := &AbstractLogger{level: ERROR}
	logger.SetWriter(func(s string) {
		fmt.Println("Error Logger:", s)
	})
	return logger
}
