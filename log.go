package main

import (
	logPkg "log"
)

func log(format string, v ...interface{}) {
	logPkg.Printf(format, v...)
}

func fatal(v ...interface{}) {
	logPkg.Fatal(v...)
}
