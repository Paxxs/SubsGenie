package MyLog

import (
	"log"
	"os"
)

type MyLogger struct {
	logger       *log.Logger
	debugEnabled bool
	infoEnabled  bool
}

func NewMyLogger() *MyLogger {
	logLevel := os.Getenv("LOG_LEVEL")
	return &MyLogger{
		logger:       log.New(os.Stdout, "", log.Ldate|log.Ltime),
		debugEnabled: logLevel == "DEBUG",
		infoEnabled:  logLevel == "DEBUG" || logLevel == "INFO",
	}
}

func (l *MyLogger) Debug(v ...interface{}) {
	if l.debugEnabled {
		l.logger.SetPrefix("DEBUG ")
		l.logger.Println(v...)
	}
}

func (l *MyLogger) Info(v ...interface{}) {
	if l.infoEnabled {
		l.logger.SetPrefix("INFO ")
		l.logger.Println(v...)
	}
}
