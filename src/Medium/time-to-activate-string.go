// https://leetcode.com/contest/weekly-contest-461/problems/minimum-time-to-activate-string/description/

func minTime(s string, order []int, k int) int {
	// instead of tracking the valid substrings, it's easier to track the clean substrings
	// store all the clean segments
	// easy to count the number of substrings in each segment
	n := len(s)
	total_possible := n * (n + 1) / 2
	if k > total_possible {
		return -1
	}
	segs := map[int]int{0: n - 1}
	cleanCount := total_possible
	for t, idx := range order {
		startSeg := 0
		endSeg := 0
		for start, end := range segs {
			if idx >= start && idx <= end {
				delete(segs, start)
				startSeg = start
				endSeg = end
				break
			}
		}
		segLen := endSeg - startSeg + 1
		currLen := segLen * (segLen + 1) / 2
		cleanCount -= currLen
		// add two new segments
		if idx > startSeg {
			leftLen := idx - startSeg
			cleanCount += leftLen * (leftLen + 1) / 2
			segs[startSeg] = idx - 1
		}
		if idx < endSeg {
			rightLen := endSeg - idx
			cleanCount += rightLen * (rightLen + 1) / 2
			segs[idx+1] = endSeg
		}
		valid := total_possible - cleanCount
		if valid >= k {
			return t
		}
	}
	return -1
}