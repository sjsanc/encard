package log

import (
	"fmt"
	"log"
	"os"

	s "github.com/sjsanc/encard/internal/styles"
)

var Logger = log.New(os.Stdout, s.Question.Render("[encard] » "), 0)

var VERBOSE bool

func Warn(format string, args ...interface{}) {
	msg := s.Warn.Render("WARN: " + fmt.Sprintf(format, args...))
	Logger.Println(msg)
}

func Error(format string, args ...interface{}) {
	msg := s.Error.Render("ERROR: " + fmt.Sprintf(format, args...))
	Logger.Println(msg)
}

func Info(format string, args ...interface{}) {
	if VERBOSE {
		msg := "INFO: " + fmt.Sprintf(format, args...)
		Logger.Println(msg)
	}
}
