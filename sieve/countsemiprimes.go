package sieve

/*
https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_semiprimes/
A prime is a positive integer X that has exactly two distinct divisors: 1 and X. The first few prime integers are 2, 3, 5, 7, 11 and 13.

A semiprime is a natural number that is the product of two (not necessarily distinct) prime numbers. The first few semiprimes are 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.

You are given two non-empty arrays P and Q, each consisting of M integers. These arrays represent queries about the number of semiprimes within specified ranges.

Query K requires you to find the number of semiprimes within the range (P[K], Q[K]), where 1 ≤ P[K] ≤ Q[K] ≤ N.

For example, consider an integer N = 26 and arrays P, Q such that:

	P[0] = 1    Q[0] = 26
	P[1] = 4    Q[1] = 10
	P[2] = 16   Q[2] = 20

The number of semiprimes within each of these ranges is as follows:

(1, 26) is 10,
(4, 10) is 4,
(16, 20) is 0.
Write a function:

class Solution { public int[] solution(int N, int[] P, int[] Q); }

that, given an integer N and two non-empty arrays P and Q consisting of M integers, returns an array consisting of M elements specifying the consecutive answers to all the queries.

For example, given an integer N = 26 and arrays P, Q such that:

	P[0] = 1    Q[0] = 26
	P[1] = 4    Q[1] = 10
	P[2] = 16   Q[2] = 20

the function should return the values [10, 4, 0], as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..50,000];
M is an integer within the range [1..30,000];
each element of arrays P and Q is an integer within the range [1..N];
P[i] ≤ Q[i].
*/

func CountSemiprimes(N int, P []int, Q []int) []int {
	// generate primes
	primes := primes(N)

	// generate subprimes
	allSemiprimes := make([]int, N+1)
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if i*j > N {
				break
			}
			if primes[i]*primes[j] == 1 {
				allSemiprimes[i*j] = 1
			}
		}
	}

	// compute prefix sums
	sum := 0
	prefixSums := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		sum += allSemiprimes[i]
		prefixSums[i] = sum
	}

	// go over P and Q and count
	M := len(P)
	counts := make([]int, M)
	for i := 0; i < M; i++ {
		counts[i] = prefixSums[Q[i]] - prefixSums[P[i]-1]
	}
	return counts
}

func primes(N int) []int {
	sieve := make([]int, N+1)

	for i := range sieve {
		sieve[i] = 1
	}

	sieve[0], sieve[1] = 0, 0
	for i := 2; i*i <= N; i++ {
		if sieve[i] == 1 {
			for k := i * i; k <= N; k += i {
				sieve[k] = 0
			}
		}
	}

	return sieve
}
