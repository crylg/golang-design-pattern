package factory_method

import (
	"log"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()
	left := 4
	var right = 2
	add := Factory(add)
	if add.Calculate(left, right) != 6 {
		t.Errorf("%d add %d error: %d", left, right, 6)
	}
	sub := Factory(sub)
	if sub.Calculate(left, right) != 2 {
		t.Errorf("%d sub %d error: %d", left, right, 2)
	}
	mul := Factory(mul)
	if mul.Calculate(left, right) != 8 {
		t.Errorf("%d mul %d error: %d", left, right, 8)
	}
	div := Factory(div)
	if div.Calculate(left, right) != 2 {
		t.Errorf("%d div %d error: %d", left, right, 2)
	}
	Factory("invalid")
}
