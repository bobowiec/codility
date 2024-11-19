package fibonacci

import "testing"

func TestFibFrog(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}, 3},
	}

	for _, test := range table {
		result := FibFrog(test.A)
		if result != test.ans {
			t.Errorf("FibFrog is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
