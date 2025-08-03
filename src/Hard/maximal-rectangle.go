// https://leetcode.com/problems/maximal-rectangle/
// chatgpt conversation: https://chatgpt.com/c/688e3c95-8058-8000-a7db-c16b9ad8c980
package hard

import (
	stk "leetcode/collections/stack"
	"slices"
)

// solve it using the largest rectangle histogram problem
func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	var m int
	if n > 0 {
		m = len(matrix[0])
	}
	// find the heights of the rectangles in each row
	// then run the max rectangle area for each row
	maxArea := 0
	heights := make([]int, m)
	for row := range n {
		heights = getRowHeights(matrix, row, heights)
		//fmt.Println(heights)
		currArea := getRowMaxRectangle(heights)
		maxArea = max(maxArea, currArea)
	}
	return maxArea
}

func getRowHeights(matrix [][]byte, row int, heights []int) []int {
	cols := len(matrix[row])
	//heights := make([]int, cols)
	for j := range cols {
		if matrix[row][j] == '0' {
			heights[j] = 0
			continue
		}
		heights[j] += 1
		// currHeight := 1
		// rowIdx := row-1
		// for rowIdx >= 0 {
		//     if matrix[rowIdx][j] == '1' {
		//         currHeight++
		//     }else{
		//         break
		//     }
		//     rowIdx--
		// }
		// heights[j] = currHeight
	}
	return heights
}

func getRowMaxRectangle(heights []int) int {
	// at each index i, find the previous idx where element is smaller than heights[i]
	//      and find the next idx where element is smaller than heights[i]
	//      then multiply by width to get the area for this rectangle
	// two stacks prevSmaller and nextSmaller -> use for getting the index
	// two arrays prev, next -> store index of prev smaller and next smaller element
	prevSmaller := stk.NewStack()
	nextSmaller := stk.NewStack()
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
