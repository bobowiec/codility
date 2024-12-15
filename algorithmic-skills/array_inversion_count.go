package algorithmicskills

/*
https://app.codility.com/programmers/trainings/4/array_inversion_count/

An array A consisting of N integers is given. An inversion is a pair of indexes (P, Q) such that P < Q and A[Q] < A[P].

Write a function:

func Solution(A []int) int

that computes the number of inversions in A, or returns −1 if it exceeds 1,000,000,000.

For example, in the following array:

 A[0] = -1 A[1] = 6 A[2] = 3
 A[3] =  4 A[4] = 7 A[5] = 4
there are four inversions:

   (1,2)  (1,3)  (1,5)  (4,5)
so the function should return 4.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
*/

func mergeAndCount(arr []int, left, mid, right int) int {
	// length of left and righ halves
	n1 := mid - left + 1
	n2 := right - mid

	// tmp arrays
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	// copying to temp arrays
	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}

	i, j, k := 0, 0, left
	inversions := 0
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
			inversions += n1 - i
		}
		k++
	}

	for ; i < n1; i++ {
		arr[k] = leftArr[i]
		k++
	}

	for ; j < n2; j++ {
		arr[k] = rightArr[j]
		k++
	}

	return inversions
}

func ArrayInversionCount(A []int) int {
	N := len(A)
	if N <= 1 {
		return 0
	}

	middle := N / 2
	leftArray := A[:middle]
	rightArray := A[middle:]

	leftInversions := ArrayInversionCount(leftArray)
	if leftInversions >= 1000000000 {
		return -1
	}

	rightInversions := ArrayInversionCount(rightArray)
	if rightInversions >= 1000000000 {
		return -1
	}

	mergedInversions := mergeAndCount(A, 0, middle-1, N-1)
	if mergedInversions >= 1000000000 {
		return -1
	}

	allInversions := leftInversions + rightInversions + mergedInversions
	if allInversions >= 1000000000 {
		return -1
	}

	return allInversions
}

func ArrayInversionCountBruteForce(A []int) int {
	N := len(A)
	if N <= 1 {
		return 0
	}
	inversions := 0
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				inversions++
				if inversions == 1000000000 {
					return -1
				}
			}
		}
	}
	return inversions
}
