package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

var Logger *log.Logger

func init() {
	logFile, err := os.OpenFile("./logs/apilogs.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	multi := io.MultiWriter(logFile, os.Stdout)
	Logger = log.New(multi, "", log.Ldate|log.Ltime)
}

func Println(message string) {
	logMessage := getMessageWithCallerInfo(message)
	Logger.Println(logMessage)
}

func Printf(format string, args ...interface{}) {
	logMessage := getMessageWithCallerInfo(format)
	Logger.Printf(logMessage, args...)
}

func Fatalln(message string) {
	logMessage := getMessageWithCallerInfo(message)
	Logger.Fatalln(logMessage)
}

func Fatalf(format string, args ...interface{}) {
	logMessage := getMessageWithCallerInfo(format)
	Logger.Fatalf(logMessage, args...)
}

func getMessageWithCallerInfo(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%v:%v-%v", file, line, message)
}
