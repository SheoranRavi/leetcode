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
// tail[i] would have the smallest number for tail ending at i+1

func lengthOfLIS2(nums []int) int {
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
