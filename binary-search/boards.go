package binarysearch

func numberOfBoards(A []int, boardSize int) int {
	boards := 0
	endOfBoard := -1
	for i := range A {
		if A[i] == 1 && endOfBoard < i {
			boards++
			endOfBoard = i + boardSize - 1
		}
	}
	return boards
}

func Boards(A []int, k int) int {
	n := len(A)
	minSize := 1
	maxSize := n
	result := -1
	for minSize <= maxSize {
		midSize := (minSize + maxSize) / 2
		boards := numberOfBoards(A, midSize)
		if boards == k {
			result = midSize
			break
		} else if boards < k {
			maxSize = midSize - 1 // shrinking board size increases # of boards
		} else {
			minSize = midSize + 1 // raising board size decreases # of boards
		}
	}
	return result
}

func Boards2(A []int, k int) int {
	n := len(A)
	minSize := 1
	maxSize := n
	result := -1
	for minSize <= maxSize {
		midSize := (minSize + maxSize) / 2
		if numberOfBoards(A, midSize) <= k {
			maxSize = midSize - 1 // shrinking board size increases # of boards
			result = midSize
		} else {
			minSize = midSize + 1 // raising board size decreases # of boards
		}
	}
	return result
}
