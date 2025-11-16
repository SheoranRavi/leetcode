// https://leetcode.com/problems/target-sum/

package medium

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2*total+1)
	}

	if total < abs(target) {
		return 0
	}

	// mark the first row
	// total is the offset in the array
	dp[0][total+nums[0]] = 1
	dp[0][total-nums[0]] = 1
	if nums[0] == 0 {
		// special case, nums[0] with either + or - gives 0 (total is just the offset)
		dp[0][total] = 2
	}

	for i := 1; i < n; i++ {
		for t := -total; t < total+1; t++ {
			if dp[i-1][total+t] > 0 {
				dp[i][total+t+nums[i]] += dp[i-1][total+t]
				dp[i][total+t-nums[i]] += dp[i-1][total+t]
			}
		}
	}
	return dp[n-1][total+target]
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
