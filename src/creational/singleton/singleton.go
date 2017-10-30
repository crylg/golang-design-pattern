package singleton

import (
	"sync"
	"log"
)

type Alone struct {
}

func GetInstance() *Alone {
	var once sync.Once
	var alone *Alone
	once.Do(func() {
		alone = &Alone{}
	})
	return alone
}

func (a Alone) Alone() {
	log.Println("alone...")
}
