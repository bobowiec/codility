package complexity

import "testing"

func TestPermMissingElem(t *testing.T) {
	var tables = []struct {
		A []int
		R int
	}{
		{[]int{2, 3, 1, 5}, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 10, 8}, 9},
	}

	for _, test := range tables {
		result := PermMissingElem(test.A)
		if result != test.R {
			t.Errorf("Missing element is invalid, got: %d, expected: %d", result, test.R)
		}
	}
}
