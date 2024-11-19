package caterpillar

import "math"

/*
https://app.codility.com/programmers/lessons/15-caterpillar_method/abs_distinct/

A non-empty array A consisting of N numbers is given. The array is sorted in non-decreasing order. The absolute distinct count of this array is the number of distinct absolute values among the elements of the array.

For example, consider array A such that:

  A[0] = -5
  A[1] = -3
  A[2] = -1
  A[3] =  0
  A[4] =  3
  A[5] =  6
The absolute distinct count of this array is 5, because there are 5 distinct absolute values among the elements of this array, namely 0, 1, 3, 5 and 6.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N numbers, returns absolute distinct count of array A.

For example, given array A such that:

  A[0] = -5
  A[1] = -3
  A[2] = -1
  A[3] =  0
  A[4] =  3
  A[5] =  6
the function should return 5, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647];
array A is sorted in non-decreasing order.
*/

func AbsDistinct(A []int) int {
	n := len(A)
	left, right := 0, n-1
	count := 0
	var lastSeen *int // pointer to the last seen abs to avoid duplicates

	for left <= right {
		leftAbs := int(math.Abs(float64(A[left])))
		rightAbs := int(math.Abs(float64(A[right])))

		if leftAbs > rightAbs {
			if lastSeen == nil || *lastSeen != leftAbs {
				count++
				lastSeen = &leftAbs
			}
			left++
		} else if leftAbs < rightAbs {
			if lastSeen == nil || *lastSeen != rightAbs {
				count++
				lastSeen = &rightAbs
			}
			right--
		} else {
			if lastSeen == nil || *lastSeen != leftAbs {
				count++
				lastSeen = &leftAbs
			}
			left++
			right--
		}
	}
	return count
}
