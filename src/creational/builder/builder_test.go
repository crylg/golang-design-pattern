package builder

import (
	"log"
	"testing"
)

func TestBuilder(t *testing.T) {
	p := new(PersonBuilder).Name("Bruce").Age(28).Build()
	if p != nil {
		log.Println(p.ToString())
	} else {
		t.Error("person builder error")
	}
}
