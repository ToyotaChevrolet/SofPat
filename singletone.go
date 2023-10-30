package main

import (
	"sync"
)

type Singleton struct {
	value int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{value: 0}
	})
	return instance
}

func main() {
	instance1 := GetInstance()
	instance1.value = 10

	instance2 := GetInstance()
}
