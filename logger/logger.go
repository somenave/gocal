package logger

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var errorLogger *log.Logger

func Init(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return file, nil
}

func Info(msg string) {
	infoLogger.Output(2, msg)
}

func Error(msg string) {
	errorLogger.Output(2, msg)
}
