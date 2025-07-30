// https://leetcode.com/problems/number-of-longest-increasing-subsequence/
// chatgpt conversation link: https://chatgpt.com/c/6887626b-ff18-8000-bc18-a40516d11854
// tag:HARD   (this is actually a hard level question)

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	count := make([]int, n)
	for i := range n {
		dp[i] = 1
		count[i] = 1
	}
	for i := range n {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					// found a new LIS, so update the count
					count[i] = count[j]
					dp[i] = dp[j] + 1
				} else if dp[j]+1 == dp[i] {
					// found another way to get same len LIS, so add the number of LIS at j
					count[i] += count[j]
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		//fmt.Println("lis for", i, ":", dp)
		//fmt.Println("count for", i, ":", count)
	}

	maxLis := 1
	for _, lis := range dp {
		maxLis = max(maxLis, lis)
	}

	res := 0
	for i, c := range count {
		// count[i] is the number of dp[i] length subsequence ending at index i
		// so, if length of subsequence ending at i is equal to max Length,
		// then add the number of LIS at i to res
		if dp[i] == maxLis {
			res += c
		}
	}
	return res
}