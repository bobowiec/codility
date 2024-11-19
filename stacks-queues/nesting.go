package stacksqueues

/*
https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/

A string S consisting of N characters is called properly nested if:

S is empty;
S has the form "(U)" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, string "(()(())())" is properly nested but string "())" isn't.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..1,000,000];
string S is made only of the characters '(' and/or ')'.
*/
type stack []string

func (st *stack) push(s string) {
	*st = append(*st, s)
}

func (st *stack) pop() string {
	if st.isEmpty() {
		return ""
	}
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *stack) isEmpty() bool {
	return len(*st) == 0
}

func Nesting(S string) int {
	st := stack{}
	for _, v := range S {
		if v == '(' {
			st.push("(")
		} else { // v == ')'
			if st.isEmpty() {
				return 0
			} else {
				st.pop()
			}
		}
	}
	if st.isEmpty() {
		return 1
	}
	return 0
}
