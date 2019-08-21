package main

import (
	"gopkg.in/qamarian-lib/str.v2"
	"gopkg.in/qamarian-mmp/rxlib.v1"
)

// NewLog () helps create a new log.
func NewLog () (*log) {
	return &log {}
}

// This data type is an emissary of the abstract data type rxlib.RxLog.
type log struct {}

func (l *log) Record (newLog string, logType byte) (error) {
	outputType := "std"
	switch logType {
		case rxlib.LogWarning: outputType = "wrn"
		case rxlib.LogError: outputType = "err"
	}
	str.PrintEtr (newLog, outputType, "rexa")
	return nil
}
