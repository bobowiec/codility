package stacksqueues

/*
https://app.codility.com/programmers/lessons/7-stacks_and_queues/stone_wall/

You are going to build a stone wall. The wall should be straight and N meters long, and its thickness should be constant; however, it should have different heights in different places. The height of the wall is specified by an array H of N positive integers. H[I] is the height of the wall from I to I+1 meters to the right of its left end. In particular, H[0] is the height of the wall's left end and H[N−1] is the height of the wall's right end.

The wall should be built of cuboid stone blocks (that is, all sides of such blocks are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

Write a function:

func Solution(H []int) int

that, given an array H of N positive integers specifying the height of the wall, returns the minimum number of blocks needed to build it.

For example, given array H containing N = 9 integers:

	H[0] = 8    H[1] = 8    H[2] = 5
	H[3] = 7    H[4] = 9    H[5] = 8
	H[6] = 7    H[7] = 4    H[8] = 8

the function should return 7. The figure shows one possible arrangement of seven blocks.
*/
type intStack []int

func (st *intStack) push(i int) {
	*st = append(*st, i)
}

func (st *intStack) isEmpty() bool {
	return len(*st) == 0
}

func (st *intStack) pop() int {
	if st.isEmpty() {
		return -1 // error
	}
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *intStack) peek() int {
	if st.isEmpty() {
		return -1 // error
	}
	top := (*st)[len(*st)-1]
	return top
}

func StoneWall(H []int) int {
	stack := intStack{}
	numBlocks := 0

	for _, height := range H {
		for !stack.isEmpty() && stack.peek() > height {
			stack.pop()
		}

		if stack.isEmpty() || stack.peek() < height {
			stack.push(height)
			numBlocks++
		}
	}

	return numBlocks
}
