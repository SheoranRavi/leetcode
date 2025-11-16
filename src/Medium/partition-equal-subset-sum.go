// https://leetcode.com/problems/partition-equal-subset-sum/
package medium

func canPartition(nums []int) bool {
	// get the total sum
	// sum1, sum2
	// subset1, subset2
	// 0/1 knapsack because element can go either in 1 or in 2
	// so take only one subset
	// sum2 = total - sum1
	total := 0
	for _, v := range nums {
		total += v
	}
	dp := make([][]int, len(nums))
	for i := range nums {
		dp[i] = make([]int, total+1)
		for j := range total + 1 {
			dp[i][j] = -1
		}
	}
	return partition(nums, 0, total, 0, dp)
}

func partition(nums []int, sum1, total int, i int, dp [][]int) bool {
	if i == len(nums) {
		return sum1 == total-sum1
	}
	if dp[i][sum1] != -1 {
		temp := dp[i][sum1]
		if temp == 0 {
			return false
		}
		return true
	}
	//include nums[i] into first subset
	take := partition(nums, sum1+nums[i], total, i+1, dp)
	//include nums[i] into second subset
	notTake := partition(nums, sum1, total, i+1, dp)
	res := take || notTake
	dp[i][sum1] = boolToInt(res)
	return res
}

func boolToInt(b bool) int {
	if b == false {
		return 0
	}
	return 1
}
