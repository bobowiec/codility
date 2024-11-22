package greedy

import "testing"

func TestMaxNonoverlappingSegments(t *testing.T) {
	var table = []struct {
		A   []int
		B   []int
		ans int
	}{
		{[]int{1, 3, 7, 9, 9}, []int{5, 6, 8, 9, 10}, 3},
	}

	for _, test := range table {
		result := MaxNonoverlappingSegments(test.A, test.B)
		if result != test.ans {
			t.Errorf("MaxNonoverlappingSegments is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
