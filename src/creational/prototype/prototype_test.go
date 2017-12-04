package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	origin := &Prototype{}
	origin.SetName("Smith")
	origin.SetSex("male")
	origin.SetAge(28)

	clone := origin.Clone()
	if &origin == &clone {
		t.Errorf("origin:%v equals clone:%v ", &origin, &clone)
	}

	clone.SetName("Bruce")
	clone.SetAge(25)
	origin.SetName("Jack")
	origin.SetAge(100)
	if &origin == &clone {
		t.Errorf("origin:%v equals clone:%v ", &origin, &clone)
	}
}
