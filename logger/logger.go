package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[37m"
	logFormat   = "2006/01/02 15:04:05"
)

var l *logger

type logger struct {
	debugEnabled bool
	debug        *log.Logger
	info         *log.Logger
	warn         *log.Logger
	error        *log.Logger
}

type logWriter struct {
	format string
	io.Writer
	sync.Mutex
}

func (w *logWriter) Write(b []byte) (n int, err error) {
	return w.Writer.Write(append([]byte(time.Now().Format(w.format)), b...))
}

// New creates a new singleton logging instance.
func New(debug bool) {
	l = &logger{
		debugEnabled: debug,
		debug: log.New(&logWriter{
			Writer: os.Stdout,
			format: logFormat,
		}, fmt.Sprintf(" [%sdbg%s] ", colorGray, colorReset), 0),
		info: log.New(&logWriter{
			Writer: os.Stdout,
			format: logFormat,
		}, " [inf] ", 0),
		warn: log.New(&logWriter{
			Writer: os.Stdout,
			format: logFormat,
		}, fmt.Sprintf(" [%swrn%s] ", colorYellow, colorReset), 0),
		error: log.New(&logWriter{
			Writer: os.Stderr,
			format: logFormat,
		}, fmt.Sprintf(" [%serr%s] ", colorRed, colorReset), 0),
	}
}

func Debug(msg string) {
	if !l.debugEnabled {
		return
	}
	l.debug.Println(msg)
}

func Debugf(format string, v ...any) {
	if !l.debugEnabled {
		return
	}
	l.debug.Printf(format, v...)
}

func Info(msg string) {
	l.info.Println(msg)
}

func Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func Warn(msg string) {
	l.warn.Println(msg)
}

func Warnf(format string, v ...any) {
	l.warn.Printf(format, v...)
}

func Error(v ...any) {
	l.error.Println(v...)
}

func Errorf(format string, v ...any) {
	l.error.Printf(format, v...)
}

func Fatal(v ...any) {
	l.error.Fatal(v...)
}

func Fatalf(format string, v ...any) {
	l.error.Fatalf(format, v)
}
