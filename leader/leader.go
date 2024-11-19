package leader

import "sort"

func slowLeader(A []int) int {
	n := len(A)
	leader := -1
	for _, k := range A {
		candidate := k
		count := 0
		for _, v := range A {
			if v == candidate {
				count++
			}
		}
		if count > n/2 {
			leader = candidate
		}
	}
	return leader
}

func fastLeader(A []int) int {
	n := len(A)
	leader := -1
	sort.Ints(A)
	candidate := A[n/2]
	count := 0
	for _, k := range A {
		if k == candidate {
			count++
		}
	}
	if count > n/2 {
		leader = candidate
	}
	return leader
}

func goldenLeader(A []int) int {
	n := len(A)
	size := 0
	value := 0 // TODO
	for _, k := range A {
		if size == 0 {
			size += 1
			value = k
		} else {
			if value != k {
				size -= 1
			} else {
				size += 1
			}
		}
	}
	candidate := -1
	if size > 0 {
		candidate = value
	}
	leader := -1
	count := 0
	for _, k := range A {
		if k == candidate {
			count++
		}
	}
	if count > n/2 {
		leader = candidate
	}
	return leader
}
