package encard

import (
	"log"
	"os"

	"github.com/sjsanc/encard/internal/styles"
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
		Logger: log.New(out, styles.Question.Bold(true).Render("[encard] ")+"» ", 0),
	}
}
