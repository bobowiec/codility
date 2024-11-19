package maxsliceproblem

/*
* https://codility.com/media/train/7-MaxSlice.pdf
 */

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func slowMaxSlice(A []int) int {
	n := len(A)
	result := 0
	for p := 0; p < n; p++ {
		for q := p; q < n; q++ {
			sum := 0
			for i := p; i < q+1; i++ {
				sum += A[i]
			}
			result = max(result, sum)
		}
	}
	return result
}

func prefixSums(A []int) []int {
	n := len(A)
	pref := make([]int, n+1)
	for k := 1; k < n+1; k++ {
		pref[k] = pref[k-1] + A[k-1]
	}
	return pref
}

func quadraticMaxSliceWithPref(A []int, pref []int) int {
	n := len(A)
	result := 0
	for p := 0; p < n; p++ {
		for q := p; q < n; q++ {
			sum := pref[q+1] - pref[p]
			result = max(result, sum)
		}
	}
	return result
}

func quadraticMaxSlice(A []int) int {
	n := len(A)
	result := 0
	for p := 0; p < n; p++ {
		sum := 0
		for q := p; q < n; q++ {
			sum += A[q]
			result = max(result, sum)
		}
	}
	return result
}

func goldenMaxSlice(A []int) int {
	maxEnding, maxSlice := 0, 0
	for _, v := range A {
		maxEnding = max(0, maxEnding+v)
		maxSlice = max(maxSlice, maxEnding)
	}
	return maxSlice
}
