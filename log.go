package main

import (
	logPkg "log"
)

type Logger struct {
	enabled bool
}

func (logger *Logger) log(format string, v ...interface{}) {
	if logger.enabled {
		logPkg.Printf(format, v...)
	}
}

func (logger *Logger) fatal(v ...interface{}) {
	logPkg.Fatal(v...)
}

var logger *Logger

func init() {
	logger = &Logger{
		enabled: true,
	}
}

func log(format string, v ...interface{}) {
	logger.log(format, v...)
}

func fatal(v ...interface{}) {
	logger.fatal(v...)
}
