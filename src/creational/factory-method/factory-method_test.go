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

	operator := NewOperatorFactory().CreateOperate("+")
	log.Printf("add result is %d\n", operator.Operator(1, 2))
	return
}
