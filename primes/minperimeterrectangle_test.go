package primes

import "testing"

func TestMinPerimeterRectangle(t *testing.T) {
	var table = []struct {
		N   int
		ans int
	}{
		{30, 22},
	}

	for _, test := range table {
		result := MinPerimeterRectangle(test.N)
		if result != test.ans {
			t.Errorf("MinPerimeterRectangle is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
