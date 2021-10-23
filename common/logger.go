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
	logger  *log.Logger
	logfile *os.File
}

func NewStdFileLogger() *StdLogger {
	return newLogger("./logs/")
}

func NewFileLogger(path string) *StdLogger {
	return newLogger(path)
}

func newLogger(path string) *StdLogger {
	filename := path + "log_" + fmt.Sprintf("%d", time.Now().UnixMilli()) + ".txt"
	lgr := log.New(os.Stdout, "", 5)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("!!!!!!!!!!! Could not create a file log !!!!!!!!!!!")
	}

	return &StdLogger{logger: lgr, logfile: file}
}

func doLog(l *StdLogger, level Level, message string) {
	msgToLog := fmt.Sprintf("[%s] %s %s\n", level, time.Now().String(), message)
	l.logger.Print(msgToLog)
	l.logfile.Write([]byte(msgToLog))
}

func logMsg(l *StdLogger, level Level, args ...interface{}) {
	doLog(l, level, fmt.Sprint(args...))
}

func logMsgf(l *StdLogger, level Level, format string, args ...interface{}) {
	doLog(l, level, fmt.Sprintf(format, args...))
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
