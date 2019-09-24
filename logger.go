package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime
)

var levelMap = map[Level]string{
	TRACE: "[TRACE] ",
	DEBUG: "[DEBUG] ",
	INFO:  "[INFO] ",
	WARN:  "[WARN] ",
	ERROR: "[ERROR] ",
}

func (l Level) String() string {
	if str, ok := levelMap[l]; ok {
		return str
	}
	return ""
}

type Logger interface {
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type StdLogger struct {
	MinLevel Level
	Lg       *log.Logger
}

// New create a new logger instance
func New(level Level, prefix string, out io.Writer, flag int) *StdLogger {
	return &StdLogger{
		MinLevel: level,
		Lg:       log.New(out, prefix, flag),
	}
}

var std = New(INFO, "", os.Stdout, Lshortfile|LstdFlags)

func sdtPrintf(level Level, format string, v ...interface{}) {
	if level >= std.MinLevel {
		std.Lg.Output(3, fmt.Sprintf(level.String()+format, v...))
	}
}

func SetMinLevel(level Level) {
	std.MinLevel = level
}

func SetOutput(out io.Writer) {
	std.Lg.SetOutput(out)
}

func SetFlags(flag int) {
	std.Lg.SetFlags(flag)
}

func SetPrefix(prefix string) {
	std.Lg.SetPrefix(prefix)
}

func Tracef(format string, v ...interface{}) {
	sdtPrintf(TRACE, format, v...)
}

func Debugf(format string, v ...interface{}) {
	sdtPrintf(DEBUG, format, v...)
}

func Infof(format string, v ...interface{}) {
	sdtPrintf(INFO, format, v...)
}

func Warnf(format string, v ...interface{}) {
	sdtPrintf(WARN, format, v...)
}

func Errorf(format string, v ...interface{}) {
	sdtPrintf(ERROR, format, v...)
}

func (l *StdLogger) logPrintf(level Level, format string, v ...interface{}) {
	if level >= l.MinLevel {
		l.Lg.Output(3, fmt.Sprintf(level.String()+format, v...))
	}
}

func (l *StdLogger) Tracef(format string, v ...interface{}) {
	l.logPrintf(TRACE, format, v...)
}

func (l *StdLogger) Debugf(format string, v ...interface{}) {
	l.logPrintf(DEBUG, format, v...)
}

func (l *StdLogger) Infof(format string, v ...interface{}) {
	l.logPrintf(INFO, format, v...)
}

func (l *StdLogger) Warnf(format string, v ...interface{}) {
	l.logPrintf(WARN, format, v...)
}

func (l *StdLogger) Errorf(format string, v ...interface{}) {
	l.logPrintf(ERROR, format, v...)
}
