package prefixsums

import "testing"

func TestPassingCars(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{0, 1, 0, 1, 1}, 5},
	}

	for _, test := range table {
		result := PassingCars(test.A)
		if result != test.ans {
			t.Errorf("Passing cars counter was incorrect, got: %d, expected: %d", result, test.ans)
		}
	}
}
