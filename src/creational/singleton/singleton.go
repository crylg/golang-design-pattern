package singleton

import (
	"log"
	"sync"
)

var a *Alone
var once sync.Once

func GetInstance() *Alone {
	once.Do(func() {
		a = &Alone{}
	})
	return a
}

type Alone struct{}

func (a Alone) Alone() {
	log.Println("a...")
}
