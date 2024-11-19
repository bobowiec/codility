package sorting

import "testing"

func TestTriangle(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{10, 2, 5, 1, 8, 10}, 1},
		{[]int{10, 50, 5, 1}, 0},
	}

	for _, test := range table {
		result := Triangle(test.A)

		if result != test.ans {
			t.Errorf("Triangle was incorrect, got: %d, expected: %d\n", result, test.ans)
		}
	}
}
