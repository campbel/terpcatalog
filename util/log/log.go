package log

import (
	"os"

	"github.com/charmbracelet/log"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr)
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(log.TextFormatter)
	logger.SetTimeFormat("15:04:05")
	logger.SetReportCaller(true)
}

func Info(msg any, keyvals ...any) {
	logger.Info(msg, keyvals...)
}

func Error(msg any, keyvals ...any) {
	logger.Error(msg, keyvals...)
}

func Fatal(msg any, keyvals ...any) {
	logger.Fatal(msg, keyvals...)
}

func FatalError(msg any, err error, keyvals ...any) {
	logger.Fatal(msg, append(keyvals, "error", err.Error())...)
}
