package log

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	DEBUG = 2
	INFO  = 1
	ERROR = 0
)

var (
	Level  = DEBUG
	red    = color.New(color.FgRed).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	debugH = cyan("debug")
	errorH = red("error")
)

func Debugf(format string, a ...interface{}) {
	if Level >= DEBUG {
		s := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s\n", debugH, s)
	}
}

func Error(err error) { Errorf(err.Error()) }

func Errorf(format string, a ...interface{}) {
	if Level >= ERROR {
		s := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s\n", errorH, s)
	}
}
