package algorithmicskills

import "testing"

func TestStrSymmetryPoint(t *testing.T) {
	var table = []struct {
		S   string
		ans int
	}{
		{"racecar", 3},
		{"x", 0},
		{"", -1},
		{"abba", -1},
	}

	for _, test := range table {
		result := StrSymmetryPoint(test.S)
		if result != test.ans {
			t.Errorf("StrSymmetryPoint is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
