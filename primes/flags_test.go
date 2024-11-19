package primes

import "testing"

func TestFlags(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}, 3},
		{[]int{0, 0, 0, 0, 0, 1, 0, 1, 0, 1}, 2},
		{[]int{1, 3, 2}, 1},
	}

	for _, test := range table {
		result := Flags(test.A)
		if result != test.ans {
			t.Errorf("Flags is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}

	for _, test := range table {
		result := SlowerFlags(test.A)
		if result != test.ans {
			t.Errorf("SlowerFlags is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
