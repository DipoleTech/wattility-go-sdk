package wattility_go_sdk

import (
	"log"
	"os"
)

type Logger interface {
	Debug()
	Print(string)
}

func NewLogger(debug bool) Logger {
	return &logging{
		debug:  debug,
		Logger: log.New(os.Stdout, "[wattility sdk]", log.Ldate|log.Ltime|log.LUTC),
	}
}

type logging struct {
	debug  bool
	Logger *log.Logger
}

func (l *logging) Debug() {
	l.debug = true
}

func (l *logging) Print(msg string) {
	if l.debug {
		l.Logger.Println(msg)
	}
}
