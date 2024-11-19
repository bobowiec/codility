package counting

import "testing"

func TestMissingInteger(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{1, 3, 6, 4, 1, 2}, 5},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -3}, 1},
	}

	for _, test := range table {
		result := MissingInteger(test.A)
		if result != test.ans {
			t.Errorf("MissingInteger was incorrect, got: %d but expected: %d", result, test.ans)
		}
	}
}
