// https://leetcode.com/problems/longest-increasing-subsequence/

// DP solution O(n^2)
package medium

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}
	for i := range n {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	maxLis := 1
	for _, lis := range dp {
		maxLis = max(lis, maxLis)
	}
	return maxLis
}

// O(n log n) solution
// we keep track of the tail of LIS at each point

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}
	for i := range n {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	maxLis := 1
	for _, lis := range dp {
		maxLis = max(lis, maxLis)
	}
	return maxLis
}
