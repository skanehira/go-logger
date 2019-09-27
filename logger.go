package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Level log level
type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

// flags is same as log pakcage
// please see `go doc log.Ldate`
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

// Logger is logger interface
type Logger interface {
	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type stdLogger struct {
	MinLevel Level
	Lg       *log.Logger
}

var out = &os.Stdout

var std = &stdLogger{
	MinLevel: INFO,
	Lg:       log.New(*out, "", Llongfile|LstdFlags),
}

func output(str string) {
	std.Lg.Output(4, str)
}

func stdPrintf(level Level, format string, v ...interface{}) {
	if level >= std.MinLevel {
		output(fmt.Sprintf(level.String()+format, v...))
	}
}

// SetMinLevel set the log min level
func SetMinLevel(level Level) {
	std.MinLevel = level
}

// SetOutput set the log output destination
func SetOutput(out io.Writer) {
	std.Lg.SetOutput(out)
}

// SetFlags set the log flags
func SetFlags(flag int) {
	std.Lg.SetFlags(flag)
}

// SetPrefix set the log prefix
func SetPrefix(prefix string) {
	std.Lg.SetPrefix(prefix)
}

// Tracef output trace log
func Tracef(format string, v ...interface{}) {
	stdPrintf(TRACE, format, v...)
}

// Debugf output debug log
func Debugf(format string, v ...interface{}) {
	stdPrintf(DEBUG, format, v...)
}

// Infof outpuut info log
func Infof(format string, v ...interface{}) {
	stdPrintf(INFO, format, v...)
}

// Warnf outpuut warnning log
func Warnf(format string, v ...interface{}) {
	stdPrintf(WARN, format, v...)
}

// Errorf output error log
func Errorf(format string, v ...interface{}) {
	stdPrintf(ERROR, format, v...)
}

	}
}

}
