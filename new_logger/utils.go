package newlogger

import "strings"

func (l *Logger) maskSensitiveWords(message string) string {
	maskedMessage := message
	for _, word := range l.wordsToMask {
		maskedMessage = strings.Replace(maskedMessage, word, strings.Repeat("*", len(word)), -1)
	}
return maskedMessage
}
