package merror

import (
	"runtime"
	"wzlog/wanzlog"

	"github.com/lytics/logrus"
)

type ErrorLevel int

const (
	LogLevel ErrorLevel = iota
	WarnLevel
	CrashLevel
)

var errorLog *log.Log

func init() {
	errorLog = log.Init("20060102.error")
}

func errorLogOutput(level ErrorLevel, a interface{}) {
	_, file, line, _ := runtime.Caller(2)
	if level == LogLevel {
		//output log
		logrus.Info(a, file, line)
		errorLog.Println(a, file, line)
	} else if level == WarnLevel {
		logrus.Warn(a, file, line)
		errorLog.Println(a, file, line)
	} else if level == CrashLevel {
		logrus.Warn(a, file, line)
		errorLog.Println(a, file, line)
	}
}

func Log(a ...interface{}) {
	errorLogOutput(LogLevel, a)
}

func Warn(a ...interface{}) {
	errorLogOutput(WarnLevel, a)
}

func Crash(a ...interface{}) {
	errorLogOutput(CrashLevel, a)
}
