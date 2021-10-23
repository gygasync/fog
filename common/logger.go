package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Level string

const (
	Info  Level = "INFO"
	Error Level = "ERROR"
	Fatal Level = "FATAL"
	Warn  Level = "WARN"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Log(level Level, args ...interface{})
	Logf(level Level, format string, args ...interface{})
}

type StdLogger struct {
	logger *log.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{logger: log.New(os.Stdout, "", 5)}
}

func logMsg(l *StdLogger, level Level, args ...interface{}) {
	msgToLog := fmt.Sprintf("[%s] %s %s", level, time.Now().String(), fmt.Sprint(args...))
	l.logger.Println(msgToLog)
}

func logMsgf(l *StdLogger, level Level, format string, args ...interface{}) {
	msgToLog := fmt.Sprintf("[%s] %s %s", level, time.Now().String(), fmt.Sprintf(format, args...))
	l.logger.Println(msgToLog)
}

func (l *StdLogger) Info(args ...interface{}) {
	logMsg(l, Info, args...)
}

func (l *StdLogger) Error(args ...interface{}) {
	logMsg(l, Error, args...)
}

func (l *StdLogger) Fatal(args ...interface{}) {
	logMsg(l, Fatal, args...)
}

func (l *StdLogger) Warn(args ...interface{}) {
	logMsg(l, Warn, args...)
}

func (l *StdLogger) Log(level Level, args ...interface{}) {
	logMsg(l, level, args...)
}

func (l *StdLogger) Infof(format string, args ...interface{}) {
	logMsgf(l, Info, format, args...)
}

func (l *StdLogger) Errorf(format string, args ...interface{}) {
	logMsgf(l, Error, format, args...)
}

func (l *StdLogger) Fatalf(format string, args ...interface{}) {
	logMsgf(l, Fatal, format, args...)
}

func (l *StdLogger) Warnf(format string, args ...interface{}) {
	logMsgf(l, Warn, format, args...)
}

func (l *StdLogger) Logf(level Level, format string, args ...interface{}) {
	logMsgf(l, level, format, args...)
}
