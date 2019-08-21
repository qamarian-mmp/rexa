package main

import (
	"gopkg.in/qamarian-lib/str.v2"
	"gopkg.in/qamarian-mmp/rxlib.v0"
)

// NewLog () helps create a new log.
func newLog () (*log) {
	return &log {}
}

// This data type is an implementation of the abstract data type (ADT) rxlib.RxLog.
type log struct {}

func (l *log) Record (newLog string, logType byte) (error) {
	outputType := "std"
	switch logType {
		case rxlib.LrtWarning: outputType = "wrn"
		case rxlib.LrtError: outputType = "err"
	}
	str.PrintEtr (newLog, outputType, "rexa")
	return nil
}
