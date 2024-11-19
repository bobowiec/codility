package primes

import "testing"

func TestPeaks(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, 3},
	}

	for _, test := range table {
		result := Peaks(test.A)
		if result != test.ans {
			t.Errorf("Peaks is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
