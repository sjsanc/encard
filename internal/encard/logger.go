package encard

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger(enabled bool) *Logger {
	out := os.Stdout
	if !enabled {
		out = nil
	}
	return &Logger{
		Logger: log.New(out, "[encard] ", log.Ltime),
	}
}
