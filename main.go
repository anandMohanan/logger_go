package main

import (
	newlogger "logger/new_logger"
)

func main() {
	logLevel := newlogger.ERROR

	wordsToMask := []string{"hello", "world"}
	log := newlogger.NewLogger(logLevel, wordsToMask)

	log.Debug("hello, the word hello should be masked. printing debug and info")
	log.Info("hi, printing info, info will be printed")
	log.Error("hi, every log will be printed")

}
