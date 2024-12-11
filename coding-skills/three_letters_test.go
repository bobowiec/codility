package codingskills

import "testing"

func TestThreeLetters(t *testing.T) {
	var table = []struct {
		A   int
		B   int
		ans string
	}{
		{5, 3, "aabaabab"}, // "abaabbaa" also ok
		{3, 3, "ababab"},   // ""aababb", "abaabb"" also ok
		{1, 4, "bbabb"},
	}

	for _, test := range table {
		result := ThreeLetters(test.A, test.B)
		if result != test.ans {
			t.Errorf("ThreeLetters is incorrect, got: %s expected: %s\n", result, test.ans)
		}
	}
}
