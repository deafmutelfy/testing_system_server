package logger

import (
	"log"
	"sync"
)

var once sync.Once

type Logger struct{}

var loggerInstance *Logger = nil

func GetInstance() *Logger {
	once.Do(func() {
		loggerInstance = &Logger{}
	})

	return loggerInstance
}

func (l *Logger) Infoln(v ...interface{}) {
	v = append([]interface{}{"[INFO]:"}, v...)
	log.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO]: "+format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR]: "+format, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	v = append([]interface{}{"[ERROR]:"}, v...)
	log.Println(v...)
}
