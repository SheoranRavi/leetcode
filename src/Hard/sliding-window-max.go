// https://leetcode.com/problems/sliding-window-maximum/
package hard

import (
	"container/list"
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	freq := make(map[int]int)
	start := 0
	end := 1
	freq[nums[0]]++

	dq := list.New()
	// use back for holding largest
	dq.PushBack(nums[0])
	for end < k {
		currMax := dq.Back().Value.(int)
		curr := nums[end]
		if curr >= currMax {
			dq.PushBack(curr)
		} else {
			// remove all elements from the front that are smaller than curr
			// dq has to be monotonically increasing
			for dq.Front().Value.(int) < curr {
				dq.Remove(dq.Front())
			}
			dq.PushFront(curr)
		}
		freq[curr]++
		end++
	}

	res := []int{dq.Back().Value.(int)}
	n := len(nums)

	for end > 0 && end < n {
		// remove start element
		freq[nums[start]]--
		currMax := dq.Back().Value.(int)
		// handle the case of the largest element being removed
		for freq[currMax] == 0 && dq.Len() > 0 {
			dq.Remove(dq.Back())
			if dq.Len() > 0 {
				currMax = dq.Back().Value.(int)
			} else {
				// reset max value if dq becomes empty
				currMax = math.MinInt
			}
		}
		start++
		curr := nums[end]
		freq[curr]++
		if curr >= currMax {
			dq.PushBack(curr)
		} else {
			for dq.Front().Value.(int) < curr {
				dq.Remove(dq.Front())
			}
			dq.PushFront(curr)
		}
		// dq back always gives the current largest in the k segment
		res = append(res, dq.Back().Value.(int))
		end++
	}
	return res
}
