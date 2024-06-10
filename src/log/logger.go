package log

import (
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

type ANSIColor string

const (
	RESET         ANSIColor = "\033[0m"
	BOLD                    = "\033[1m"
	ITALIC                  = "\033[3m"
	UNDERLINE               = "\033[4m"
	INVERSE                 = "\033[7m"
	STRIKETHROUGH           = "\033[9m"
	BLACK                   = "\033[30m"
	RED                     = "\033[31m"
	GREEN                   = "\033[32m"
	YELLOW                  = "\033[33m"
	BLUE                    = "\033[34m"
	MAGENTA                 = "\033[35m"
	CYAN                    = "\033[36m"
	WHITE                   = "\033[37m"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

func (l LogLevel) String() string {
	return [...]string{"INFO", "WARN", "ERROR"}[l]
}

type Logger struct {
	mu     sync.Mutex
	output *log.Logger
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{
		output: log.New(out, "", log.LstdFlags),
	}
}

func (l *Logger) log(color ANSIColor, text string, withTimestamp bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if withTimestamp {
		timestamp := time.Now().Format(time.RFC3339)
		logText := fmt.Sprintf("[%q]: %s", timestamp, text)
		l.output.Println(colorize(color, logText))
	} else {
		l.output.Println(colorize(color, text))
	}
}

func (l *Logger) Info(text string) {
	l.log(BLUE, text, false)
}

func (l *Logger) Warn(text string) {
	l.log(YELLOW, text, false)
}

func (l *Logger) Error(text string) {
	l.log(RED, text, false)
}

func colorize(color ANSIColor, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, RESET)
}
