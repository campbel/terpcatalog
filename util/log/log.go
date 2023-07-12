package log

import (
	"os"

	"github.com/campbel/terpcatalog/util/config"
	"github.com/charmbracelet/log"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr)
	if config.IsDevelopment() {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetLevel(log.InfoLevel)
	}
	logger.SetFormatter(log.TextFormatter)
	logger.SetTimeFormat("15:04:05")
	logger.SetReportCaller(true)
	logger.SetCallerOffset(1)
}

func Debug(msg any, keyvals ...any) {
	logger.Debug(msg, keyvals...)
}

func Info(msg any, keyvals ...any) {
	logger.Info(msg, keyvals...)
}

func Error(msg any, err error, keyvals ...any) {
	logger.Error(msg, append(keyvals, "error", err.Error())...)
}

func Fatal(msg any, keyvals ...any) {
	logger.Fatal(msg, keyvals...)
}

func FatalError(msg any, err error, keyvals ...any) {
	logger.Fatal(msg, append(keyvals, "error", err.Error())...)
}
