package logging

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type Level int

var CurrentLoggingLevel Level
var ColorLogLevelLabelOnly = true

var greenPrinter, yellowPrinter, redPrinter *color.Color

const (
	BlankLevel Level = 10
	InfoLevel  Level = 3
	WarnLevel  Level = 2
	DebugLevel Level = 1
)

func init() {
	CurrentLoggingLevel = InfoLevel
	greenPrinter = color.New(color.FgGreen)
	yellowPrinter = color.New(color.FgYellow)
	redPrinter = color.New(color.FgRed)
}

//SetLevel allows settings of the level of logging
func SetLevel(loggingLevel Level) {
	CurrentLoggingLevel = loggingLevel
}

//Info outputs log line to console with green color text
func Info(format string, a ...interface{}) (n int, err error) {
	if CurrentLoggingLevel > InfoLevel {
		return
	}

	printFunc := greenPrinter.Printf
	strPrintFunc := greenPrinter.Sprintf
	if ColorLogLevelLabelOnly {
		printFunc = fmt.Printf
		strPrintFunc = fmt.Sprintf
	}

	return printFunc("%s %s %s\n", getTimeString(), color.GreenString("[INFO]"), strPrintFunc(format, a...))
}

//Warn outputs log line to console with yellow color text
func Warn(format string, a ...interface{}) (n int, err error) {
	if CurrentLoggingLevel > WarnLevel {
		return
	}

	printFunc := yellowPrinter.Printf
	strPrintFunc := yellowPrinter.Sprintf
	if ColorLogLevelLabelOnly {
		printFunc = fmt.Printf
		strPrintFunc = fmt.Sprintf
	}

	return printFunc("%s %s %s\n", getTimeString(), color.YellowString("[WARN]"), strPrintFunc(format, a...))
}

//Debug outputs log line to console with yellow color text
func Debug(format string, a ...interface{}) (n int, err error) {
	if CurrentLoggingLevel > DebugLevel {
		return
	}

	printFunc := yellowPrinter.Printf
	strPrintFunc := yellowPrinter.Sprintf
	if ColorLogLevelLabelOnly {
		printFunc = fmt.Printf
		strPrintFunc = fmt.Sprintf
	}

	return printFunc("%s %s %s\n", getTimeString(), color.YellowString("[DEBUG]"), strPrintFunc(format, a...))
}

//Error outputs log line to console with red color text
func Error(format string, a ...interface{}) (n int, err error) {
	printFunc := redPrinter.Printf
	strPrintFunc := redPrinter.Sprintf
	if ColorLogLevelLabelOnly {
		printFunc = fmt.Printf
		strPrintFunc = fmt.Sprintf
	}

	return printFunc("%s %s %s\n", getTimeString(), color.RedString("[ERROR]"), strPrintFunc(format, a...))
}

//ErrorAndExit outputs log line to console with red color text and exits
func Fatal(format string, a ...interface{}) {
	Error(format, a...)
	os.Exit(1)
}

func createCallbackLabel(skip int) string {
	function, _, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("(%s):%d", runtime.FuncForPC(function).Name(), line)
}

//getTimeString gets formatted string to timestamp log and console output
func getTimeString() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
