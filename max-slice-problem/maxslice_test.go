package maxsliceproblem

import "testing"

func TestMaxSlice(t *testing.T) {
	var table = []struct {
		A   []int
		ans int
	}{
		{[]int{5, -7, 3, 5, -2, 4, -1}, 10},
	}

	for _, test := range table {
		result := slowMaxSlice(test.A)
		if result != test.ans {
			t.Errorf("SlowMaxSlice is incorrect, got: %d expected: %d\n", result, test.ans)
		}

		pref := prefixSums(test.A)
		result = quadraticMaxSliceWithPref(test.A, pref)
		if result != test.ans {
			t.Errorf("QuadraticMaxSliceWithPrefSum is incorrect, got: %d expected: %d\n", result, test.ans)
		}

		result = quadraticMaxSlice(test.A)
		if result != test.ans {
			t.Errorf("QuadraticMaxSlice is incorrect, got: %d expected: %d\n", result, test.ans)
		}

		result = goldenMaxSlice(test.A)
		if result != test.ans {
			t.Errorf("GoldenMaxSlice is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
