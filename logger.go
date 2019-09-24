package logger

import (
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

type Logger struct {
	MinLevel Level
	Lg       *log.Logger
}

// New create a new logger instance
func New(level Level, prefix string, out io.Writer, flag int) *Logger {
	return &Logger{
		MinLevel: level,
		Lg:       log.New(out, prefix, flag),
	}
}

var std *Logger = New(INFO, "", os.Stdout, log.Lshortfile|log.LstdFlags)

func sdtPrintf(level Level, format string, v ...interface{}) {
	if std.MinLevel >= level {
		std.Lg.Printf(level.String()+format+"\n", v)
	}
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
