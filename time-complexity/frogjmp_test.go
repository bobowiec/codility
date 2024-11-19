package complexity

import "testing"

func TestCountFrogJumps(t *testing.T) {
	var tests = []struct {
		X int
		Y int
		D int
		R int
	}{
		{10, 85, 30, 3},
		{1, 9, 1, 8},
		{10, 10, 1, 0},
	}

	for _, test := range tests {
		result := CountFrogJumps(test.X, test.Y, test.D)
		if result != test.R {
			t.Errorf("Counter of frog jumps incorrect, got: %d, expected: %d", result, test.R)
		}
	}
}
