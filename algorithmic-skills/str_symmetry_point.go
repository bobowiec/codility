package algorithmicskills

/*
https://app.codility.com/programmers/trainings/4/str_symmetry_point/

Write a function:

func Solution(S string) int

that, given a string S, returns the index (counting from 0) of a character such that the part of the string to the left of that character is a reversal of the part of the string to its right. The function should return −1 if no such index exists.

Note: reversing an empty string (i.e. a string whose length is zero) gives an empty string.

For example, given a string:

"racecar"

the function should return 3, because the substring to the left of the character "e" at index 3 is "rac", and the one to the right is "car".

Given a string:

"x"

the function should return 0, because both substrings are empty.

Write an efficient algorithm for the following assumptions:

the length of string S is within the range [0..2,000,000].
*/

func StrSymmetryPoint(S string) int {
	N := len(S)

	// skip even length strings
	if N%2 == 0 {
		return -1
	}

	// special case 1 character
	if N == 1 {
		return 0
	}

	left := 0
	right := N - 1
	for left < right {
		if S[left] != S[right] {
			return -1
		}
		left++
		right--
	}
	if left == right {
		return N / 2
	}
	return -1
}
