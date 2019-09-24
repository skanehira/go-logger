package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

var (
	stdout io.Writer = os.Stdout
)

func errorf(tb testing.TB, want, got interface{}) {
	tb.Helper()
	tb.Errorf("want = %+v, got = %+v", want, got)
}

func makeWantLog(prefix, level, format, args string) string {
	return prefix + level + fmt.Sprintf(format, args) + "\n"
}

func TestNew(t *testing.T) {
	level := INFO
	out := os.Stdout
	prefix := "[gorilla]"
	flag := log.LstdFlags

	logger := New(level, prefix, out, flag)

	if logger.MinLevel != level {
		errorf(t, level, logger.MinLevel)
	}

	if logger.Lg.Flags() != flag {
		errorf(t, flag, logger.Lg.Flags())
	}

	if logger.Lg.Prefix() != prefix {
		errorf(t, prefix, logger.Lg.Prefix())
	}

	lgout := logger.Lg.Writer().(*os.File)
	if lgout.Name() != out.Name() {
		errorf(t, out.Name(), lgout.Name())
	}
}

func TestLevel(t *testing.T) {
	var GORILLA Level = 5
	var tests = []struct {
		name string
		lv   Level
		want string
	}{
		{"trace string", TRACE, "[TRACE] "},
		{"debug string", DEBUG, "[DEBUG] "},
		{"info string", INFO, "[INFO] "},
		{"warn string", WARN, "[WARN] "},
		{"error string", ERROR, "[ERROR] "},
		{"undefined level", GORILLA, ""},
	}

	for _, test := range tests {
		t.Run("[Level to string] "+test.name, func(t *testing.T) {
			if test.lv.String() != test.want {
				errorf(t, test.want, test.lv.String())
			}
		})
	}
}

func TestSetFunc(t *testing.T) {
	testLogger := New(INFO, "", os.Stdout, log.Lshortfile)
	oldLogger := std
	std = testLogger
	defer func() { std = oldLogger }()

	SetMinLevel(TRACE)
	if std.MinLevel != TRACE {
		errorf(t, std.MinLevel.String(), TRACE.String())
	}

	out := stdout.(*os.File)
	SetOutput(out)
	lgout := std.Lg.Writer().(*os.File)
	if lgout.Name() != out.Name() {
		errorf(t, out.Name(), out.Name())
	}

	prefix := "gorilla"
	SetPrefix(prefix)
	if std.Lg.Prefix() != prefix {
		errorf(t, prefix, std.Lg.Prefix())
	}

	flag := log.Ldate
	SetFlags(flag)
	if std.Lg.Flags() != flag {
		errorf(t, flag, std.Lg.Flags())
	}
}

func TestStdPrintf(t *testing.T) {
	var buf bytes.Buffer
	prefix := "[test] "
	std = New(DEBUG, prefix, &buf, 0)

	format := "I am %s"
	args := "gorilla"

	tests := []struct {
		name      string
		level     Level
		printFunc func(format string, v ...interface{})
		excepted  string
	}{
		{"global std trace log", TRACE, Tracef, ""},
		{"global std debug log", DEBUG, Debugf, makeWantLog(prefix, DEBUG.String(), format, args)},
		{"global std info log", INFO, Infof, makeWantLog(prefix, INFO.String(), format, args)},
		{"global std warn log", WARN, Warnf, makeWantLog(prefix, WARN.String(), format, args)},
		{"global std error log", ERROR, Errorf, makeWantLog(prefix, ERROR.String(), format, args)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			excepted := test.excepted
			test.printFunc(format, args)
			if buf.String() != excepted {
				errorf(t, excepted, buf.String())
			}
			buf.Reset()
		})
	}
}

func TestLoggerPrintf(t *testing.T) {
	var buf bytes.Buffer
	prefix := "[test] "
	std := New(DEBUG, prefix, &buf, 0)

	format := "I am %s"
	args := "gorilla"

	tests := []struct {
		name      string
		level     Level
		printFunc func(format string, v ...interface{})
		excepted  string
	}{
		{"logger print trace log", TRACE, std.Tracef, ""},
		{"logger print debug log", DEBUG, std.Debugf, makeWantLog(prefix, DEBUG.String(), format, args)},
		{"logger print info log", INFO, std.Infof, makeWantLog(prefix, INFO.String(), format, args)},
		{"logger print warn log", WARN, std.Warnf, makeWantLog(prefix, WARN.String(), format, args)},
		{"logger print error log", ERROR, std.Errorf, makeWantLog(prefix, ERROR.String(), format, args)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			excepted := test.excepted
			test.printFunc(format, args)
			if buf.String() != excepted {
				errorf(t, excepted, buf.String())
			}
			buf.Reset()
		})
	}
}
