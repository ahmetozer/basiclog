package basiclog

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

type logFunc func(format string, v ...any)

var (
	Debug, Info, Error, Fatal logFunc
	flag                                = log.Ldate | log.Ltime
	out                       io.Writer = os.Stdout
)

func SetFlag(f int) {
	flag = f
}

func SetOut(o io.Writer) {
	out = o
}

func Init() {
	var dummyLog logFunc = func(format string, v ...any) {}

	var err error
	if logFile := os.Getenv("LOG_FILE"); logFile != "" {
		out, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Log file cannot open: %v %s", logFile, err)
		}
	}

	Debug = log.New(out, "DEBUG: ", flag).Printf
	Info = log.New(out, "INFO: ", flag).Printf
	Error = log.New(out, "ERROR: ", flag).Printf
	Fatal = log.New(out, "FATAL: ", flag).Fatalf

	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "DEBUG":

	case "INFO":
		Debug = dummyLog
	case "ERROR":
		Debug, Info = dummyLog, dummyLog
	case "FATAL":
		Debug, Info, Error = dummyLog, dummyLog, dummyLog
	default:
		Debug = dummyLog
	}
}

func ErrNil(f logFunc, err error, errs ...error) bool {

	if len(errs) > 0 {
		err = errors.Join(append([]error{err}, errs...)...)
	}
	if err != nil {
		f("%v", err)
		return false
	}
	return true

}
