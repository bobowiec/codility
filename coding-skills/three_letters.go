package codingskills

/*
https://app.codility.com/programmers/trainings/5/three_letters/

Write a function Solution that, given two integers A and B, returns a string containing exactly A letters 'a' and exactly B letters 'b' with no three consecutive letters being the same (in other words, neither "aaa" nor "bbb" may occur in the returned string).

Examples:

1. Given A = 5 and B = 3, your function may return "aabaabab". Note that "abaabbaa" would also be a correct answer. Your function may return any correct answer.

2. Given A = 3 and B = 3, your function should return "ababab", "aababb", "abaabb" or any of several other strings.

3. Given A = 1 and B = 4, your function should return "bbabb", which is the only correct answer in this case.

Assume that:

A and B are integers within the range [0..100];
at least one solution exists for the given A and B.
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
*/
func ThreeLetters(A, B int) string {
	var maxChar, minChar string
	max, min := 0, 0

	if A >= B {
		max = A
		maxChar = "a"
		min = B
		minChar = "b"
	} else {
		max = B
		maxChar = "b"
		min = A
		minChar = "a"
	}

	result := ""
	for max != min {
		if max > 1 {
			result += maxChar
			result += maxChar
			max -= 2
		} else {
			result += maxChar
			max--
		}

		if min > 0 {
			result += minChar
			min--
		}
	}

	for max != 0 && min != 0 {
		if max > 0 {
			result += maxChar
			max--
		}
		if min > 0 {
			result += minChar
			min--
		}
	}

	return result
}
