package complexity

import "testing"

func TestTapeEqulibrum(t *testing.T) {
	var cases = []struct {
		A   []int
		ans int
	}{
		{[]int{3, 1, 2, 4, 3}, 1},
		{[]int{10, 3, 1, 2, 4, 3}, 3},
		{[]int{10, 3, 10}, 3},
		{[]int{-2, -3, -4, -1}, 0},
		{[]int{4, 5, 1, 8}, 0},
		{[]int{-10, -20, -30, -40, 100}, 20},
		{[]int{-10, -5, -3, -4, -5}, 3},
		{[]int{1, 2}, 1},
	}

	for _, test := range cases {
		result := TapeEquilibrium(test.A)
		if result != test.ans {
			t.Errorf("TapeEquilibrium was incorrect, got: %d but expected: %d", result, test.ans)
		}
	}
}
