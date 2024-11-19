package primes

import "testing"

func TestCountFactors(t *testing.T) {
	var table = []struct {
		n   int
		ans int
	}{
		{2, 2},
		{3, 2},
		{4, 3},
		{5, 2},
		{6, 4},
		{7, 2},
		{8, 4},
		{9, 3},
		{24, 8},
	}

	for _, test := range table {
		result := CountFactors(test.n)
		if result != test.ans {
			t.Errorf("CountFactors is incorrect, n=%d got: %d expected: %d\n", test.n, result, test.ans)
		}
	}
}
