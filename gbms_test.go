package gbms

import "testing"

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}
