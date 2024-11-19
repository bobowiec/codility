package prefixsums

import (
	"testing"
)

func TestCountDiv(t *testing.T) {
	var table = []struct {
		A   int
		B   int
		K   int
		ans int
	}{
		{6, 11, 2, 3},
		{11, 345, 17, 20},
		{0, 0, 11, 1},
		{10, 10, 5, 1},
		{10, 10, 7, 0},
		{10, 10, 20, 0},
		{2, 15, 3, 5},
	}

	for _, test := range table {
		result := CountDiv(test.A, test.B, test.K)
		if result != test.ans {
			t.Errorf("Counting div was incorrect, got: %d, expected: %d", result, test.ans)
		}
	}
}
