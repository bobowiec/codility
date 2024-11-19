package sieve

import "testing"

func TestCountSemiprimes(t *testing.T) {
	var table = []struct {
		N       int
		P       []int
		Q       []int
		answers []int
	}{
		{26, []int{1, 4, 16}, []int{26, 10, 20}, []int{10, 4, 0}},
	}

	for _, test := range table {
		results := CountSemiprimes(test.N, test.P, test.Q)
		for i, r := range results {
			if r != test.answers[i] {
				t.Errorf("CountSemiprimes is incorrect, got: %d expected: %d\n", r, test.answers[i])
			}
		}
	}
}
