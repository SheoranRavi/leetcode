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