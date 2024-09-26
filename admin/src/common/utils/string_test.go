package utils

import "testing"

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandStr(16, true))
		t.Log(RandStr(16, false))
	}
}
