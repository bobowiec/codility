package stacksqueues

import (
	"testing"
)

func TestFish(t *testing.T) {
	var table = []struct {
		A   []int
		B   []int
		ans int
	}{
		{[]int{4, 3, 2, 1, 5}, []int{0, 1, 0, 0, 0}, 2},
		{[]int{0, 1}, []int{1, 1}, 2},
	}

	for _, test := range table {
		result := Fish(test.A, test.B)
		if result != test.ans {
			t.Errorf("Fish is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
