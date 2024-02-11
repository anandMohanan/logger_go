package newlogger

import (
	"fmt"
	"time"
)

func (l *Logger) log(level LogLevel, logMessage string, args ...interface{}) {

	if level > l.level {
		return
	}
	message := fmt.Sprintf(logMessage, args...)
	message = l.maskSensitiveWords(message)
	if len(message) > maxMessageLength {
		message = message[:maxMessageLength] + " [log trimmed]"
	}
	timestamp := time.Now().Format("2024-02-28 15:04:05")
	logLevel := [...]string{"INFO", "DEBUG", "ERROR"}[level]
	logMessageFmt := fmt.Sprintf("[%s] [%s] %s\n", timestamp, logLevel, message)
	fmt.Print(logMessageFmt)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}
