package singleton

import "sync"

type singleton struct {
}

var instance *singleton
var once sync.Once

func NewSingleton() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
