package iterations

import "testing"

func TestSolution(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{9, 2},
		{529, 4},
		{20, 1},
		{15, 0},
		{32, 0},
		{1041, 5},
	}

	for _, test := range tests {
		binGap := Solution(test.n)
		if binGap != test.want {
			t.Errorf("Max binary gaps of %d was incorrect, got: %d, want: %d.", test.n, binGap, test.want)
		}
	}
}
