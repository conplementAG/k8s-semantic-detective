package logging

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	colorable "github.com/mattn/go-colorable"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	file                 *os.File
	consoleLogger        *logrus.Logger
	fileLogger           *logrus.Logger
)

func AdjustVerbosity(logVerbose bool) {
	if logVerbose {
		consoleLogger.SetLevel(logrus.DebugLevel)
	} else {
		consoleLogger.SetLevel(logrus.InfoLevel)
	}
}

func Initialize(logFileName string) {
	consoleLogger = logrus.New()
	fileLogger = logrus.New()

	fileLogger.SetLevel(logrus.DebugLevel)
	consoleLogger.SetLevel(logrus.InfoLevel)

	consoleLogger.SetFormatter(&prefixed.TextFormatter{
		ForceColors:     true,
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	fileLogger.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   true,
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	fileLog := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    25,
		MaxBackups: 10,
		LocalTime:  true,
	}

	consoleLogger.SetOutput(colorable.NewColorableStdout())
	fileLogger.SetOutput(fileLog)
}

func SilenceConsole() {
	consoleLogger.SetLevel(logrus.ErrorLevel)
}

func Dispose() {
	file.Close()
}

func Info(text string) {
	consoleLogger.Info(text)
	fileLogger.Info(text)
}

func Infof(text string, v ...interface{}) {
	consoleLogger.Infof(text, v...)
	fileLogger.Infof(text, v...)
}

func Debug(text string) {
	consoleLogger.Debug(text)
	fileLogger.Debug(text)
}

func Debugf(text string, v ...interface{}) {
	consoleLogger.Debugf(text, v...)
	fileLogger.Debugf(text, v...)
}

func Fatal(text string) {
	consoleLogger.Fatal(text)
	fileLogger.Fatal(text)
}

func Fatalf(text string, v ...interface{}) {
	consoleLogger.Fatalf(text, v...)
	fileLogger.Fatalf(text, v...)
}

func Error(text string) {
	consoleLogger.Error(text)
	fileLogger.Error(text)
}

func Errorf(text string, v ...interface{}) {
	consoleLogger.Errorf(text, v...)
	fileLogger.Errorf(text, v...)
}

func LogStartSpinner(text string) *spinner.Spinner {
	spinner := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spinner.Prefix = text
	spinner.Color("green", "bold")
	spinner.Start()
	return spinner
}

func LogStopSpinner(spinner *spinner.Spinner) {
	spinner.Stop()
}
