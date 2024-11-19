package arrays

import "testing"

func TestOddOccurrencesInArray(t *testing.T) {
	var tests = []struct {
		a   []int
		ans int
	}{
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
		{[]int{1, 1, 1, 1, 2, 2, 3}, 3},
	}

	for _, test := range tests {
		result := OddOccurrences(test.a)
		if result != test.ans {
			t.Errorf("Odd occurrences in array incorrect, got: %d, expected: %d", result, test.ans)
		}
	}

}
