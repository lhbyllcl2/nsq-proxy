package tool

import (
	"nsq-proxy/internal/module/logger"
	"runtime/debug"
)

func PanicHandlerForLog() {
	if err := recover(); err != nil {
		logger.Errorf("recover panic: %v\r\n========\r\n%s", err, string(debug.Stack()))
	}
}
