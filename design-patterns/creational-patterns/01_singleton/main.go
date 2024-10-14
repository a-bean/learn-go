package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}
}
