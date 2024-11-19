package leader

/*
https://app.codility.com/programmers/lessons/8-leader/equi_leader/

A non-empty array A consisting of N integers is given.

The leader of this array is the value that occurs in more than half of the elements of A.

An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

For example, given array A such that:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
we can find two equi leaders:

0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
The goal is to count the number of equi leaders.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

For example, given:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
the function should return 2, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].
*/

type intStack []int

func (st *intStack) push(v int) {
	*st = append(*st, v)
}

func (st *intStack) pop() int {
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *intStack) peek() int {
	top := (*st)[len(*st)-1]
	return top
}

func (st *intStack) isEmpty() bool {
	return len(*st) == 0
}

func findLeader(A []int) int {
	n := len(A)
	stack := intStack{}
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
	return leader
}

func countValue(A []int, value int) int {
	counter := 0
	for _, v := range A {
		if v == value {
			counter++
		}
	}
	return counter
}

func EquiLeader(A []int) int {
	leader := findLeader(A)
	equiLeader := 0
	if leader == -1 {
		return equiLeader
	}

	n := len(A)
	leaderCount := countValue(A, leader)
	leaderCountSoFar := 0
	for i, v := range A {
		if v == leader {
			leaderCountSoFar++
		}
		if leaderCountSoFar*2 > i+1 && (leaderCount-leaderCountSoFar)*2 > n-i-1 {
			equiLeader++
		}
	}
	return equiLeader
}
