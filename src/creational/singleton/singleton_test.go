package singleton

import (
	"testing"
)

func generateAloneInstance(e interface{}, t *testing.T) {
	for i := 0; i < 100; i++ {
		a := new(Alone)
		if e != a {
			t.Errorf("singleton pattern error: %v and %v", e, a)
		}
	}
}

func TestSingleton(t *testing.T) {
	generateAloneInstance(GetInstance(), t)
}
