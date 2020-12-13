package log

import (
	"github.com/ggdream/gopack/tools/color"
	"os"
)

func (l *Logger) Error(format string, args ...interface{}) {
	l.WrapLog(LOG_LEVEL_ERROR, color.ColorRed, format, args...)
}

func (l *Logger) Warning(format string, args ...interface{}) {
	l.WrapLog(LOG_LEVEL_WARNING, color.ColorYellow, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.WrapLog(LOG_LEVEL_INFO, color.ColorGreen, format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.WrapLog(LOG_LEVEL_DEBUG, color.ColorPurple, format, args...)
}

func New() *Logger {
	logger := new(Logger)

	logger.SetLevel(LOG_LEVEL_DEBUG)
	logger.SetWriter(os.Stdout)
	logger.SetColorSupport(true)
	logger.OpenTimer()

	return logger
}
