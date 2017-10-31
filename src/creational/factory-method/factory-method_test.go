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

	operate := NewOperateFactory().CreateOperate("+")
	log.Printf("add result is %d\n", operate.Operate(1, 2))
	return
}
