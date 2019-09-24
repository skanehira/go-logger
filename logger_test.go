package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

var (
	stdout io.Writer = os.Stdout
)

func TestNew(t *testing.T) {
	level := INFO
	out := os.Stdout
	prefix := "[gorilla]"
	flag := log.LstdFlags

	logger := New(level, prefix, out, flag)

	if logger.MinLevel != level {
		t.Errorf("excepted:%s, got:%s\n", level, logger.MinLevel)
	}

	if logger.Lg.Flags() != flag {
		t.Errorf("excepted:%d, got:%d\n", flag, logger.Lg.Flags())
	}

	if logger.Lg.Prefix() != prefix {
		t.Errorf("excepted:%s, got:%s\n", prefix, logger.Lg.Prefix())
	}

	lgout := logger.Lg.Writer().(*os.File)
	if lgout.Name() != out.Name() {
		t.Errorf("excepted:%s, got:%s\n", out.Name(), lgout.Name())
	}
}

func TestLevel(t *testing.T) {
	var tests = []struct {
		lv       Level
		excepted string
	}{
		{TRACE, "[TRACE] "},
		{DEBUG, "[DEBUG] "},
		{INFO, "[INFO] "},
		{WARN, "[WARN] "},
		{ERROR, "[ERROR] "},
	}

	for _, test := range tests {
		if test.lv.String() != test.excepted {
			t.Errorf("excepted:%s, got:%s\n", test.excepted, test.lv.String())
		}
	}
}

func TestSetFunc(t *testing.T) {
	testLogger := New(INFO, "", os.Stdout, log.Lshortfile)
	oldLogger := std
	std = testLogger
	defer func() { std = oldLogger }()

	SetMinLevel(INFO)
	if std.MinLevel != INFO {
		t.Errorf("excepted:%s, got:%s\n", std.MinLevel.String(), INFO.String())
	}

	out := stdout.(*os.File)
	SetOutput(out)
	lgout := std.Lg.Writer().(*os.File)
	if lgout.Name() != out.Name() {
		t.Errorf("excepted:%s, got:%s\n", out.Name(), out.Name())
	}

	prefix := "gorilla"
	SetPrefix(prefix)
	if std.Lg.Prefix() != prefix {
		t.Errorf("excepted:%s, got:%s\n", prefix, std.Lg.Prefix())
	}

	flag := log.Ldate
	SetFlags(flag)
	if std.Lg.Flags() != flag {
		t.Errorf("excepted:%d, got:%d\n", flag, std.Lg.Flags())
	}
}

type testbuf struct {
	buf []byte
}

func (t *testbuf) Write(p []byte) (n int, err error) {
	t.buf = p
	return len(p), nil
}

func (t *testbuf) String() string {
	return string(t.buf)
}

func TestStdPrintf(t *testing.T) {
	var buf testbuf
	prefix := "[test] "
	std = New(DEBUG, prefix, &buf, 0)

	format := "I am %s"
	args := "gorilla"

	makeExcepted := func(level, format, args string) string {
		return prefix + level + fmt.Sprintf(format, args) + "\n"
	}

	tests := []struct {
		level     Level
		printFunc func(format string, v ...interface{})
		excepted  string
	}{
		{TRACE, Tracef, ""},
		{DEBUG, Debugf, makeExcepted(DEBUG.String(), format, args)},
		{INFO, Infof, makeExcepted(INFO.String(), format, args)},
		{WARN, Warnf, makeExcepted(WARN.String(), format, args)},
		{ERROR, Errorf, makeExcepted(ERROR.String(), format, args)},
	}

	for _, test := range tests {
		excepted := test.excepted
		test.printFunc(format, args)
		if buf.String() != excepted {
			t.Errorf("excepted:%s, got:%s", excepted, buf.String())
		}
	}
}
