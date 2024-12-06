package codingskills

import "testing"

func TestParityDegree(t *testing.T) {
	var table = []struct {
		N   int
		ans int
	}{
		{24, 3},
	}

	for _, test := range table {
		result := ParityDegree(test.N)
		if result != test.ans {
			t.Errorf("ParityDegree is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
