package leader

/*
https://app.codility.com/programmers/lessons/8-leader/dominator/

An array A consisting of N integers is given. The dominator of array A is the value that occurs in more than half of the elements of A.

For example, consider array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

Write a function

func Solution(A []int) int

that, given an array A consisting of N integers, returns index of any element of array A in which the dominator of A occurs. The function should return −1 if array A does not have a dominator.

For example, given array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
the function may return 0, 2, 4, 6 or 7, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
*/

type stack []int

func (st *stack) push(v int) {
	*st = append(*st, v)
}

func (st *stack) pop() int {
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *stack) peek() int {
	top := (*st)[len(*st)-1]
	return top
}

func (st *stack) isEmpty() bool {
	return len(*st) == 0
}

func Dominator(A []int) int {
	n := len(A)
	stack := stack{}
	for _, v := range A {
		if stack.isEmpty() {
			stack.push(v)
		} else {
			if stack.peek() != v {
				stack.pop()
			} else {
				stack.push(v)
			}
		}
	}
	candidate := -1
	if !stack.isEmpty() {
		candidate = stack.pop()
	}
	count := 0
	for _, v := range A {
		if v == candidate {
			count++
		}
	}
	leader := -1
	if count > n/2 {
		leader = candidate
	}

	if leader != -1 {
		for i, v := range A {
			if leader == v {
				return i
			}
		}
	}
	return -1
}
