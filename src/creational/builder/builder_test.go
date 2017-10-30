package builder

import (
	"testing"
	"log"
)

func TestBuilder(t *testing.T) {
	p := new(PersonBuilder).SetName("Bruce").SetAge(28).Build()
	log.Println(p)
}
