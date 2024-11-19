package fibonacci

/*
https://app.codility.com/programmers/lessons/13-fibonacci_numbers/fib_frog/

The Fibonacci sequence is defined using the following recursive formula:

    F(0) = 0
    F(1) = 1
    F(M) = F(M - 1) + F(M - 2) if M >= 2
A small frog wants to get to the other side of a river. The frog is initially located at one bank of the river (position −1) and wants to get to the other bank (position N). The frog can jump over any distance F(K), where F(K) is the K-th Fibonacci number. Luckily, there are many leaves on the river, and the frog can jump between the leaves, but only in the direction of the bank at position N.

The leaves on the river are represented in an array A consisting of N integers. Consecutive elements of array A represent consecutive positions from 0 to N − 1 on the river. Array A contains only 0s and/or 1s:

0 represents a position without a leaf;
1 represents a position containing a leaf.
The goal is to count the minimum number of jumps in which the frog can get to the other side of the river (from position −1 to position N). The frog can jump between positions −1 and N (the banks of the river) and every position containing a leaf.

For example, consider array A such that:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
The frog can make three jumps of length F(5) = 5, F(3) = 2 and F(5) = 5.

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the minimum number of jumps by which the frog can get to the other side of the river. If the frog cannot reach the other side of the river, the function should return −1.

For example, given:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer that can have one of the following values: 0, 1.
*/

type Queue []int // queue

// Enqueue adds a value to the queue.
func (q *Queue) enqueue(v int) {
	*q = append(*q, v)
}

// Dequeue removes and returns the first value from the queue.
func (q *Queue) dequeue() int {
	if q.isEmpty() {
		return -1 // sentinel
	}
	dequeued := (*q)[0]
	*q = (*q)[1:]
	return dequeued
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func generateFibo(N int) []int {
	fib := []int{1, 2}
	for {
		next := fib[len(fib)-1] + fib[len(fib)-2]
		if next > N {
			break
		}
		fib = append(fib, next)
	}
	return fib
}

func FibFrog(A []int) int {
	n := len(A)
	fib := generateFibo(n + 1)

	// Early return: check if we can jump directly from -1 to the other side
	for _, f := range fib {
		if f == n+1 { // Direct jump possible
			return 1
		}
	}

	// BFS setup
	position := -1
	queue := Queue{}
	queue.enqueue(position)
	jumps := Queue{}
	jumps.enqueue(0)
	visited := make([]bool, n+1)

	for !queue.isEmpty() {
		position = queue.dequeue()
		step := jumps.dequeue()

		for _, f := range fib {
			nextPosition := position + f

			// if the next position is exactly the other side of the river, return the number of jumps
			if nextPosition == n {
				return step + 1
			}

			if nextPosition >= 0 && nextPosition < n && A[nextPosition] == 1 && !visited[nextPosition] {
				visited[nextPosition] = true
				queue.enqueue(nextPosition)
				jumps.enqueue(step + 1)
			}
		}
	}
	return -1
}
