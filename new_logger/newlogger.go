package newlogger

const maxMessageLength = 300

type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	ERROR
)

type Logger struct {
	level       LogLevel
	wordsToMask []string
}

func NewLogger(level LogLevel, wordsToMask []string) *Logger {
	var loglevel LogLevel
	switch level {
	case INFO:
		loglevel = INFO

	case DEBUG:
		loglevel = DEBUG

	case ERROR:
		loglevel = ERROR

	}

	return &Logger{
		level:       loglevel,
		wordsToMask: wordsToMask,
	}

}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}
