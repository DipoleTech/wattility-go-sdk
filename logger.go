package wattility_go_sdk

import (
	"log"
	"os"
)

type logger interface {
	Debug()
	Print(string)
}

func NewLogger(debug bool) logger {
	return &logging{
		debug:  debug,
		Logger: log.New(os.Stdout, "", log.LUTC),
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
