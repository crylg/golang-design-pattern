package factory_method

import (
	"log"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	add := NewFactory().CreateOperate("+")
	log.Printf("add result is %d\n", add.Calculate(1, 2))
	sub := NewFactory().CreateOperate("-")
	log.Printf("sub result is %d\n", sub.Calculate(1, 2))
	return
}
