package main

import (
	"errors"

	"github.com/masa213f/go-util"
)

func main() {
	err := errors.New("my error")

	jsonLogger := util.NewLogger(util.LogEncodingJSON, true)
	jsonLogger.Info("info messag (json)", "key", "value", "key2", 123)
	jsonLogger.WithName("LoggerName").Error(err, "error message (json)")
	jsonLogger.V(util.LogLevelDebug).Info("debug message (json)")
	// {"lv":"info","ts":"2021-10-15T23:15:20.798186+0900","msg":"info messag (json)","key":"value","key2":123}
	// {"lv":"error","ts":"2021-10-15T23:15:20.798251+0900","logger":"LoggerName","msg":"error message (json)","error":"my error"}
	// {"lv":"debug","ts":"2021-10-15T23:15:20.798257+0900","msg":"debug message (json)"}

	textLogger := util.NewLogger(util.LogEncodingConsole, false)
	textLogger.Info("info messag (console)", "key", "value", "key2", 123)
	textLogger.WithName("LoggerName").Error(err, "error message (console)")
	textLogger.V(util.LogLevelDebug).Info("debug message (console)")
	// 2021-10-15T23:15:20.798262+0900 info    info messag (console)   {"key": "value", "key2": 123}
	// 2021-10-15T23:15:20.798269+0900 error   LoggerName      error message (console) {"error": "my error"}
}
