package logger

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

var Log *Logger

func init() {
	Log = &Logger{
		infoLogger:  log.New(os.Stdout, "", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		file = filepath.Base(file)
		l.infoLogger.Printf("INFO: %s:%d: %s", file, line, msg)
	} else {
		l.infoLogger.Println("INFO: " + msg)
	}
}

func (l *Logger) Error(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		file = filepath.Base(file)
		l.errorLogger.Printf("ERROR: %s:%d: %s", file, line, msg)
	} else {
		l.errorLogger.Println("ERROR: " + msg)
	}
}
