// https://leetcode.com/problems/russian-doll-envelopes/

// tle solution uses dynamic programming and has a time complexity O(n^2)
// this is same as longest increasing subsequence

package hard

import (
	"cmp"
	"slices"
)

func maxEnvelopes(envelopes [][]int) int {
	// [2,3] [5,4] [6,4] [6,7] [7, 5]
	slices.SortFunc(envelopes, func(a, b []int) int {
		if a[0] != b[0] {
			return cmp.Compare(a[0], b[0])
		}
		return cmp.Compare(a[1], b[1])
	})

	n := len(envelopes)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}
	for i := range n {
		for j := i + 1; j < n; j++ {
			if envelopes[j][0] > envelopes[i][0] && envelopes[j][1] > envelopes[i][1] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	maxLen := 1
	for _, num := range dp {
		maxLen = max(num, maxLen)
	}
	return maxLen
}

// we reduce the problem to 1d LIS by sorting it ascending by width, and descending by height if widths are equal
// then find the LIS in the heights array
// heights are sorted in descending order to avoid counting the cases where width are equal
func maxEnvelopes(envelopes [][]int) int {
	slices.SortFunc(envelopes, func(a, b []int) int {
		if a[0] != b[0] {
			return cmp.Compare(a[0], b[0])
		}
		return cmp.Compare(b[1], a[1])
	})

	var heights []int
	for _, env := range envelopes {
		heights = append(heights, env[1])
	}
	maxEnv := lengthOfLIS(heights)
	return maxEnv
}

func lengthOfLIS(nums []int) int {
	tails := make([]int, 0)
	tails = append(tails, nums[0])
	for _, num := range nums {
		// num either goes to end of tails
		if num > tails[len(tails)-1] {
			tails = append(tails, num)
		} else {
			idx := binarySearch(tails, num)
			tails[idx] = num
		}
	}
	return len(tails)
}

func binarySearch(tails []int, target int) int {
	// find the first number in tails that is >= target
	l := 0
	r := len(tails)
	mid := (l + r) / 2
	for l < r {
		if tails[mid] > target {
			r = mid
		} else if tails[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
		mid = (l + r) / 2
	}
	return mid
}
