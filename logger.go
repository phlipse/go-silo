package silo

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const (
	DEBUG = 1 << iota
	INFO
	WARN
	ERROR
)

type Logger interface {
	Debug()
	Info()
	Warn()
	Error()
}

type logger struct {
	sync.Mutex
	output io.Writer
	level  int
	buf    []byte
}

var (
	loggerInstance *logger
	loggerInit     sync.Once
)

func Init(o io.Writer, l int) *logger {
	loggerInit.Do(func() {
		loggerInstance = &logger{
			output: o,
			level:  l,
		}
	})

	return loggerInstance
}

func Get() *logger {
	if loggerInstance == nil {
		return &logger{
			output: os.Stdout,
			level:  ERROR,
		}
	}

	return loggerInstance
}

func (l *logger) write(m string, level int) {
	switch level {
	case DEBUG:
		m = fmt.Sprintf("%s [DEBUG] %s", time.Now().Format(time.RFC3339), m)
	case INFO:
		m = fmt.Sprintf("%s [INFO] %s", time.Now().Format(time.RFC3339), m)
	case WARN:
		m = fmt.Sprintf("%s [WARN] %s", time.Now().Format(time.RFC3339), m)
	case ERROR:
		m = fmt.Sprintf("%s [ERROR] %s", time.Now().Format(time.RFC3339), m)
	default:
		return
	}

	fmt.Fprintln(l.output, m)
}

func (l *logger) Debug(m string, v ...interface{}) {
	if !(DEBUG < l.level) {
		m = fmt.Sprintf(m, v...)
		l.write(m, DEBUG)
	}
}

func (l *logger) Info(m string, v ...interface{}) {
	if !(INFO < l.level) {
		m = fmt.Sprintf(m, v...)
		l.write(m, INFO)
	}
}

func (l *logger) Warn(m string, v ...interface{}) {
	if !(WARN < l.level) {
		m = fmt.Sprintf(m, v...)
		l.write(m, WARN)
	}
}

func (l *logger) Error(m string, v ...interface{}) {
	if !(ERROR < l.level) {
		m = fmt.Sprintf(m, v...)
		l.write(m, ERROR)
	}
}
