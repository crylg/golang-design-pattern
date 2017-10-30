package prototype

import (
	"testing"
	"log"
)

func TestPrototype(t *testing.T) {
	origin := &Prototype{}
	origin.SetName("Smith")
	origin.SetSex("male")
	origin.SetAge(28)

	clone := origin.Clone()
	log.Printf("origin origin:%v, clone:%v ", origin, clone)

	clone.SetName("Bruce")
	clone.SetAge(25)
	log.Printf("modify origin:%v, clone:%v ", origin, clone)
}
