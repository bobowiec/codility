package sorting

import "testing"

func TestDiscIntersections(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 5, 2, 1, 4, 0}, 11},
	}

	for _, test := range table {
		result := DiscIntersect(test.A)
		if result != test.ans {
			t.Errorf("NumberOfDiscIntersections is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
