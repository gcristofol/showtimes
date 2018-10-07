package main

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

func NewLog(logpath string) {
	println("LogFile: " + logpath)
	if logpath == "<undefined>" {
		Log = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		file, err := os.Create(logpath)
		if err != nil {
			panic(err)
		}
		Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	}
}
