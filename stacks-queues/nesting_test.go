package stacksqueues

import (
	"testing"
)

func TestNesting(t *testing.T) {
	var table = []struct {
		S   string
		ans int
	}{
		{"(()(())())", 1},
		{"())", 0},
	}

	for _, test := range table {
		result := Nesting(test.S)
		if result != test.ans {
			t.Errorf("Nesting is incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
