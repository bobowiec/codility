package stacksqueues

import (
	"testing"
)

func TestBrackets(t *testing.T) {
	var table = []struct {
		S   string
		ans int
	}{
		{"{[()()]}", 1},
		{"([)()]", 0},
		{"{{{{", 0},
	}

	for _, test := range table {
		result := Brackets(test.S)
		if result != test.ans {
			t.Errorf("Brackets are incorrect, got: %d expected: %d\n", result, test.ans)
		}
	}
}
