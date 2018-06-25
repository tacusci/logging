package logging

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type Level int

var currentLoggingLevel Level

const (
	InfoLevel  Level = 2
	DebugLevel Level = 1
)

func init() {
	currentLoggingLevel = InfoLevel
}

func createCallbackLabel() string {
	function, _, line, _ := runtime.Caller(2)
	return fmt.Sprintf("(%s):%d", runtime.FuncForPC(function).Name(), line)
}

//SetLevel allows settings of the level of logging
func SetLevel(loggingLevel Level) {
	currentLoggingLevel = loggingLevel
}

//ColoredOutput helper to make it easy to logout with date time stamp
func ColoredOutput(colorPrinter *color.Color, stringToPrint string) {
	colorPrinter.Printf(stringToPrint)
}

func GreenOutput(stringToPrint string) {
	green := color.New(color.FgGreen)
	green.Printf(stringToPrint)
}

func YellowOutput(stringToPrint string) {
	yellow := color.New(color.FgYellow)
	yellow.Printf(stringToPrint)
}

func RedOutput(stringToPrint string) {
	red := color.New(color.FgRed)
	red.Printf(stringToPrint)
}

//Info outputs log line to console with green color text
func Info(stringToPrint string) {
	if currentLoggingLevel <= InfoLevel {
		GreenOutput(fmt.Sprintf("%s: INFO %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//InfoNnl outputs log line to console with green color text without newline
func InfoNnl(stringToPrint string) {
	if currentLoggingLevel <= InfoLevel {
		GreenOutput(fmt.Sprintf("%s: INFO %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//Debug outputs log line to console with yellow color text
func Debug(stringToPrint string) {
	if currentLoggingLevel <= DebugLevel {
		YellowOutput(fmt.Sprintf("%s: DEBUG %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//DebugNnl outputs log line to console with yellow color text without newline
func DebugNnl(stringToPrint string) {
	if currentLoggingLevel <= DebugLevel {
		YellowOutput(fmt.Sprintf("%s: DEBUG %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//Error outputs log line to console with red color text
func Error(stringToPrint string) {
	RedOutput(fmt.Sprintf("%s: ERROR %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
}

//ErrorNnl outputs log line to console with red color text without newline
func ErrorNnl(stringToPrint string) {
	RedOutput(fmt.Sprintf("%s: ERROR %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
}

//ErrorAndExit outputs log line to console with red color text and exits
func ErrorAndExit(stringToPrint string) {
	Error(stringToPrint)
	os.Exit(1)
}

//ErrorAndExitNnl outputs the log line to the console with red color text with no newline and exits
func ErrorAndExitNnl(stringToPrint string) {
	ErrorNnl(stringToPrint)
	os.Exit(1)
}

//GetTimeString gets formatted string to timestamp log and console output
func GetTimeString() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
