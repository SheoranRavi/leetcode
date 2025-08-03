// https://leetcode.com/problems/largest-rectangle-in-histogram/description/

// simple solution that calculates the longest sequence of height at least h for each height in the 0 to maxHeight + 1 range
// Gives TLE

func largestRectangleArea(heights []int) int {
	maxHeight := 0
	for _, val := range heights {
		maxHeight = max(maxHeight, val)
	}
	maxArea := 0
	for h := range maxHeight + 1 {
		if h == 0 {
			continue
		}
		// find the longest sequence of height h in the array
		maxSeq := 0
		currSeq := 0
		for _, val := range heights {
			if val >= h {
				currSeq++
				maxSeq = max(maxSeq, currSeq)
			} else {
				currSeq = 0
			}
		}
		currArea := maxSeq * h
		maxArea = max(maxArea, currArea)
	}
	return maxArea
}

// Using monotonic stacks, get the previous smaller and next smaller element for each height
func largestRectangleArea(heights []int) int {

	// at each index i, find the previous idx where element is smaller than heights[i]
	//      and find the next idx where element is smaller than heights[i]
	//      then multiply by width to get the area for this rectangle
	// two stacks prevSmaller and nextSmaller -> use for getting the index
	// two arrays prev, next -> store index of prev smaller and next smaller element
	prevSmaller := NewStack()
	nextSmaller := NewStack()
	prev := make([]int, 0)
	next := make([]int, 0)
	for idx, v := range heights {
		for {
			top := prevSmaller.Peek()
			if top == nil {
				prev = append(prev, -1)
				break
			}
			topIdx := top.(int)
			if heights[topIdx] < v {
				prev = append(prev, topIdx)
				break
			}
			prevSmaller.Pop()
		}
		prevSmaller.Push(idx)
	}

	for idx := len(heights) - 1; idx >= 0; idx-- {
		v := heights[idx]
		for {
			top := nextSmaller.Peek()
			if top == nil {
				next = append(next, -1)
				break
			}
			topIdx := top.(int)
			if heights[topIdx] < v {
				next = append(next, topIdx)
				break
			}
			nextSmaller.Pop()
		}
		nextSmaller.Push(idx)
	}
	slices.Reverse(next)

	maxArea := 0
	//fmt.Println(prev)
	//fmt.Println(next)
	for idx, h := range heights {
		// find left and right ranges for this height
		left := prev[idx]
		right := next[idx]
		if right == -1 {
			right = len(heights)
		}
		width := right - left - 1
		area := width * h
		//fmt.Println("width for height", h, width)
		maxArea = max(maxArea, area)
	}
	return maxArea

}

type Stack struct {
	stack []int
	n     int
}

func NewStack() Stack {
	return Stack{
		stack: make([]int, 0),
		n:     0,
	}
}

func (this *Stack) Pop() any {
	if this.n == 0 {
		return nil
	}
	old := this.stack
	item := old[this.n-1:]
	this.stack = old[:this.n-1]
	this.n--
	return item[0]
}

func (this *Stack) Peek() any {
	if this.n == 0 {
		return nil
	}
	return this.stack[this.n-1]
}

func (this *Stack) Push(item int) {
	this.stack = append(this.stack, item)
	this.n++
}