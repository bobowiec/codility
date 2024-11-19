package stacksqueues

/*
https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/

A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..200,000];
string S is made only of the following characters: '(', '{', '[', ']', '}' and/or ')'.
*/

func Brackets(S string) int {
	if len(S) == 0 {
		return 1
	}

	var stack []string
	bracketsMap := map[string]string{")": "(", "}": "{", "]": "["}
	for _, ch := range S {
		s := string(ch)
		switch s {
		case "(", "{", "[":
			stack = append(stack, s)
		case ")", "}", "]":
			if len(stack) > 0 {
				n := len(stack) - 1
				top := stack[n]
				if top != bracketsMap[s] {
					return 0
				}
				stack = stack[:n]
			} else {
				return 0
			}
		}
	}

	if len(stack) == 0 {
		return 1
	} else {
		return 0
	}
}
