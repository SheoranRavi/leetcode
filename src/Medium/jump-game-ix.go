// https://leetcode.com/problems/jump-game-ix/
package medium

func maxValue(nums []int) []int {
	// separate the array into sets of connected components
	// the max within each component is the ans for all elements within that component
	n := len(nums)
	if n == 0 {
		return nil
	}
	maxLeft := make([]int, n)
	currMax := nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > currMax {
			currMax = nums[i]
		}
		maxLeft[i] = currMax
	}

	minRight := make([]int, n)
	currMin := nums[n-1]
	for i := n - 1; i >= 0; i-- {
		if nums[i] < currMin {
			currMin = nums[i]
		}
		minRight[i] = currMin
	}

	segStart := 0
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		// if maxLeft at this point is <= minRight in the next part, then it means no elements in left part can
		// jump to the right part. Hence this marks a connected component
		if i == n-1 || maxLeft[i] <= minRight[i+1] {
			for j := segStart; j <= i; j++ {
				ans[j] = maxLeft[i]
			}
			segStart = i + 1
		}
	}
	return ans
}
