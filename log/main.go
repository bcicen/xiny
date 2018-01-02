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
	debugH = color.CyanString("debug")
	infoH  = color.GreenString("info")
	errorH = color.RedString("error")
)

func Debugf(format string, a ...interface{}) {
	if Level >= DEBUG {
		s := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s\n", debugH, s)
	}
}

func Infof(format string, a ...interface{}) {
	if Level >= INFO {
		s := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s\n", infoH, s)
	}
}

func Error(err error) { Errorf(err.Error()) }

func Errorf(format string, a ...interface{}) {
	if Level >= ERROR {
		s := fmt.Sprintf(format, a...)
		fmt.Printf("%s %s\n", errorH, s)
	}
}
