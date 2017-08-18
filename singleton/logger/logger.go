package logger

import (
	"fmt"
	"sync"
)

type logger struct {
	name string
}

func (l logger) Log(level LogLevel, msg string) string {
	return fmt.Sprintf("%v :: %v %s\n", l.name, level, msg)
}

var instance *logger
var once sync.Once

func NewLogger(name string) *logger {
	once.Do(func() {
		instance = new(logger)
		instance.name = name
	})
	return instance
}
