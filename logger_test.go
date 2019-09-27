package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
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

func makeWantLog(file, prefix, level, format, args string) string {
	return prefix + file + level + fmt.Sprintf(format, args) + "\n"
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

func TestSetFuncs(t *testing.T) {
	t.Run("set properties to global logger instance", func(t *testing.T) {
		testLogger := &stdLogger{MinLevel: INFO, Lg: log.New(os.Stdout, "", Llongfile|LstdFlags)}
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

		flag := Ldate
		SetFlags(flag)
		if std.Lg.Flags() != flag {
			errorf(t, flag, std.Lg.Flags())
		}
	})
}

func TestPrintLogLevel(t *testing.T) {
	var buf bytes.Buffer
	prefix := "[test] "
	std = &stdLogger{MinLevel: DEBUG, Lg: log.New(&buf, prefix, Lshortfile)}

	format := "I am %s"
	args := "gorilla"

	tests := []struct {
		name      string
		level     Level
		printFunc func(format string, v ...interface{})
		excepted  string
	}{
		{"global std trace log", TRACE, Tracef, ""},
		{"global std debug log", DEBUG, Debugf, makeWantLog("logger_test.go:107: ", prefix, DEBUG.String(), format, args)},
		{"global std info log", INFO, Infof, makeWantLog("logger_test.go:107: ", prefix, INFO.String(), format, args)},
		{"global std warn log", WARN, Warnf, makeWantLog("logger_test.go:107: ", prefix, WARN.String(), format, args)},
		{"global std error log", ERROR, Errorf, makeWantLog("logger_test.go:107: ", prefix, ERROR.String(), format, args)},
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

func TestPrintJSON(t *testing.T) {
	var buf bytes.Buffer
	std = &stdLogger{MinLevel: DEBUG, Lg: log.New(&buf, "", 0)}

	tests := []struct {
		kind string
		name string
		data interface{}
		want string
	}{
		{"success", "success to print json", &struct{ Name string }{"gorilla"}, "{\"Name\":\"gorilla\"}\n"},
		{"failed", "failed to printj json", math.NaN(), "json: unsupported value: NaN\n"},
	}

	for _, test := range tests {
		PrintToJSON(test.data)

		if test.want != buf.String() {
			errorf(t, test.want, buf.String())
		}
		buf.Reset()
	}
}

func TestPrintStruct(t *testing.T) {
	v := struct{ Name string }{"gorilla"}

	var buf bytes.Buffer
	std = &stdLogger{MinLevel: DEBUG, Lg: log.New(&buf, "", 0)}

	PrintStruct(v)

	want := "struct { Name string }{Name:\"gorilla\"}\n"

	if buf.String() != want {
		errorf(t, want, buf.String())
	}

	buf.Reset()
}
