package main

import "fmt"

type Logger interface {
	Log()
}

type InfoLog struct {
	Message string
}

type ErrorLog struct {
	Message string
	Error   error
}

func (a *InfoLog) Log() {
	fmt.Printf("Info log: %s\n", a.Message)
}

func (a *ErrorLog) Log() {
	fmt.Printf("Error log: %s, err: %+v\n", a.Message, a.Error)
}
