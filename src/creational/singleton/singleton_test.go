package singleton

import "testing"

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Errorf("singleton pattern error with %v and %v", ins1, ins2)
	}
}
