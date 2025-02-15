package main

import (
	"errors"
	"fmt"
	e "github.com/pkg/errors"
	"runtime"
)

func main1() {
	err := errors.New("错误")
	SetError1(err)
}
func SetError1(err error) {
	msg := e.WithStack(err)
	fmt.Printf("%+v\n", msg)
}

func SetError(err error) {
	var msg = make([]byte, 1024)
	n := runtime.Stack(msg, false)
	fmt.Println(string(msg[:n]))
}