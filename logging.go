package logging

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type Level int

var CurrentLoggingLevel Level
var LoggingOutputReciever chan string
var ColorLogLevelLabelOnly = false
var OutputLogLevelFlag = true
var OutputPath bool = true
var OutputDateTime bool = true
var OutputArrowSuffix bool = true

const (
	BlankLevel Level = 10
	InfoLevel  Level = 2
	DebugLevel Level = 1
)

func init() {
	CurrentLoggingLevel = InfoLevel
}

func createCallbackLabel() string {
	function, _, line, _ := runtime.Caller(2)
	return fmt.Sprintf("(%s):%d", runtime.FuncForPC(function).Name(), line)
}

//SetLevel allows settings of the level of logging
func SetLevel(loggingLevel Level) {
	CurrentLoggingLevel = loggingLevel
}

//ColoredOutput helper to make it easy to logout with date time stamp
func ColoredOutput(colorPrinter *color.Color, stringToPrint string) {
	if LoggingOutputReciever != nil {
		LoggingOutputReciever <- stringToPrint
	}
	colorPrinter.Printf(stringToPrint)
	color.Unset()
}

func GreenOutput(stringToPrint string) {
	if LoggingOutputReciever != nil {
		LoggingOutputReciever <- stringToPrint
	}
	green := color.New(color.FgGreen)
	green.Printf(stringToPrint)
	color.Unset()
}

func YellowOutput(stringToPrint string) {
	if LoggingOutputReciever != nil {
		LoggingOutputReciever <- stringToPrint
	}
	yellow := color.New(color.FgYellow)
	yellow.Printf(stringToPrint)
	color.Unset()
}

func RedOutput(stringToPrint string) {
	if LoggingOutputReciever != nil {
		LoggingOutputReciever <- stringToPrint
	}
	red := color.New(color.FgRed)
	red.Printf(stringToPrint)
	color.Unset()
}

func WhiteOutput(stringToPrint string) {
	if LoggingOutputReciever != nil {
		LoggingOutputReciever <- stringToPrint
	}
	white := color.New(color.FgWhite)
	white.Printf(stringToPrint)
	color.Unset()
}

//Info outputs log line to console with green color text
func Info(stringToPrint string) {
	if CurrentLoggingLevel <= InfoLevel {
		if ColorLogLevelLabelOnly == false {
			GreenOutput(createOutputString(stringToPrint, "INFO", true))
			// GreenOutput(fmt.Sprintf("%s: INFO %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
		} else {
			WhiteOutput(fmt.Sprintf("%s:", GetTimeString()))
			GreenOutput(" INFO ")
			WhiteOutput(fmt.Sprintf("%s -> %s\n", createCallbackLabel(), stringToPrint))
		}
	}
}

//InfoNnl outputs log line to console with green color text without newline
func InfoNnl(stringToPrint string) {
	if CurrentLoggingLevel <= InfoLevel {
		if ColorLogLevelLabelOnly == false {
			GreenOutput(createOutputString(stringToPrint, "INFO", false))
			// GreenOutput(fmt.Sprintf("%s: INFO %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
		} else {
			WhiteOutput(fmt.Sprintf("%s:", GetTimeString()))
			GreenOutput(" INFO ")
			WhiteOutput(fmt.Sprintf("%s -> %s", createCallbackLabel(), stringToPrint))
		}
	}
}

//Info outputs log line to console with green color text
func InfoNoColor(stringToPrint string) {
	if CurrentLoggingLevel <= InfoLevel {
		WhiteOutput(createOutputString(stringToPrint, "INFO", true))
		// GreenOutput(fmt.Sprintf("%s: INFO %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//InfoNnl outputs log line to console with green color text without newline
func InfoNnlNoColor(stringToPrint string) {
	if CurrentLoggingLevel <= InfoLevel {
		WhiteOutput(createOutputString(stringToPrint, "INFO", false))
		// GreenOutput(fmt.Sprintf("%s: INFO %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//Debug outputs log line to console with yellow color text
func Debug(stringToPrint string) {
	if CurrentLoggingLevel <= DebugLevel {
		if ColorLogLevelLabelOnly == false {
			YellowOutput(createOutputString(stringToPrint, "DEBUG", true))
			//YellowOutput(fmt.Sprintf("%s: DEBUG %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
		} else {
			WhiteOutput(fmt.Sprintf("%s:", GetTimeString()))
			YellowOutput(" DEBUG ")
			WhiteOutput(fmt.Sprintf("%s -> %s\n", createCallbackLabel(), stringToPrint))
		}
	}
}

//DebugNnl outputs log line to console with yellow color text without newline
func DebugNnl(stringToPrint string) {
	if CurrentLoggingLevel <= DebugLevel {
		if ColorLogLevelLabelOnly == false {
			YellowOutput(createOutputString(stringToPrint, "DEBUG", false))
			// YellowOutput(fmt.Sprintf("%s: DEBUG %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
		} else {
			WhiteOutput(fmt.Sprintf("%s:", GetTimeString()))
			YellowOutput(" DEBUG ")
			WhiteOutput(fmt.Sprintf("%s -> %s", createCallbackLabel(), stringToPrint))
		}
	}
}

//Debug outputs log line to console with yellow color text
func DebugNoColor(stringToPrint string) {
	if CurrentLoggingLevel <= DebugLevel {
		WhiteOutput(createOutputString(stringToPrint, "DEBUG", true))
		//YellowOutput(fmt.Sprintf("%s: DEBUG %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//DebugNnl outputs log line to console with yellow color text without newline
func DebugNnlNoColor(stringToPrint string) {
	if CurrentLoggingLevel <= DebugLevel {
		WhiteOutput(createOutputString(stringToPrint, "DEBUG", false))
		// YellowOutput(fmt.Sprintf("%s: DEBUG %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
	}
}

//Error outputs log line to console with red color text
func Error(stringToPrint string) {
	if ColorLogLevelLabelOnly == false {
		RedOutput(createOutputString(stringToPrint, "ERROR", true))
		// RedOutput(fmt.Sprintf("%s: ERROR %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
	} else {
		WhiteOutput(fmt.Sprintf("%s:", GetTimeString()))
		RedOutput(" ERROR ")
		WhiteOutput(fmt.Sprintf("%s -> %s\n", createCallbackLabel(), stringToPrint))
	}
}

//ErrorNnl outputs log line to console with red color text without newline
func ErrorNnl(stringToPrint string) {
	if ColorLogLevelLabelOnly == false {
		RedOutput(createOutputString(stringToPrint, "ERROR", true))
		// RedOutput(fmt.Sprintf("%s: ERROR %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
	} else {
		WhiteOutput(fmt.Sprintf("%s:", GetTimeString()))
		RedOutput(" ERROR ")
		WhiteOutput(fmt.Sprintf("%s -> %s", createCallbackLabel(), stringToPrint))
	}
}

//Error outputs log line to console with red color text
func ErrorNoColor(stringToPrint string) {
	WhiteOutput(createOutputString(stringToPrint, "ERROR", true))
	// RedOutput(fmt.Sprintf("%s: ERROR %s -> %s\n", GetTimeString(), createCallbackLabel(), stringToPrint))
}

//ErrorNnl outputs log line to console with red color text without newline
func ErrorNnlNoColor(stringToPrint string) {
	WhiteOutput(createOutputString(stringToPrint, "ERROR", true))
	// RedOutput(fmt.Sprintf("%s: ERROR %s -> %s", GetTimeString(), createCallbackLabel(), stringToPrint))
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

//ErrorAndExit outputs log line to console with red color text and exits
func ErrorAndExitNoColor(stringToPrint string) {
	ErrorNoColor(stringToPrint)
	os.Exit(1)
}

//ErrorAndExitNnl outputs the log line to the console with red color text with no newline and exits
func ErrorAndExitNnlNoColor(stringToPrint string) {
	ErrorNnlNoColor(stringToPrint)
	os.Exit(1)
}

//GetTimeString gets formatted string to timestamp log and console output
func GetTimeString() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func createOutputString(stp string, lvl string, nl bool) string {
	data := make([]byte, 0)
	sb := bytes.NewBuffer(data)
	if OutputDateTime {
		sb.WriteString(fmt.Sprintf("%s: ", GetTimeString()))
	}
	if OutputLogLevelFlag {
		sb.WriteString(lvl)
	}
	if OutputPath {
		sb.WriteString(fmt.Sprintf(" %s", createCallbackLabel()))
	}
	if OutputArrowSuffix {
		sb.WriteString(fmt.Sprintf(" -> %s", stp))
	} else {
		sb.WriteString(stp)
	}
	if nl {
		sb.WriteString("\n")
	}
	return sb.String()
}
