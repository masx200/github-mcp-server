package streamablehttp

import (
	"log"

	"github.com/mark3labs/mcp-go/util"
)

type logAdapter struct {
	logger *log.Logger
}

// 实现 Infof
func (l *logAdapter) Infof(format string, v ...any) {
	l.logger.Printf("[INFO] "+format, v...)
}

// 实现 Errorf
func (l *logAdapter) Errorf(format string, v ...any) {
	l.logger.Printf("[ERROR] "+format, v...)
}

// 构造函数
func NewLoggerAdapter(l *log.Logger) util.Logger {
	return &logAdapter{logger: l}
}