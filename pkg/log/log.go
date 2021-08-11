package log

import (
	"github.com/fatih/color"
)

func Info(format string, a ...interface{}) {
	println(color.HiGreenString(format, a...))
}

func Error(format string, a ...interface{}) {
	println(color.HiRedString(format, a...))
}
