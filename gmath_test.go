package gobasic

import "testing"

func TestRound(t *testing.T) {
	var toTest = map[float64]int64{7.5: 8, 8: 8, -7.5: -8, 0: 0, 0.00001: 0, 10000.49: 10000}

	for in, out := range toTest {

		if x := Round(in); x != out {
			t.Errorf("Round(%v) = %v, want %v", in, x, out)
		}
	}
}
