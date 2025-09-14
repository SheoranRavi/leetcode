// https://leetcode.com/contest/weekly-contest-467/problems/number-of-stable-subsequences/submissions/1770108254/
package hard

const MOD = 1_000_000_007

func countStableSubsequences(nums []int) int {
	n := len(nums)
	dp := make([][][]int, n)
	for i := range n {
		dp[i] = make([][]int, 2)
		for p := range 2 {
			dp[i][p] = make([]int, 2)
		}
	}
	// dp[i][p][k] is the number of subsequences upto i with parity p and streak k
	firstParity := parity(nums[0])
	dp[0][firstParity][0] = 1

	for i := range n {
		if i == 0 {
			continue
		}
		currParity := parity(nums[i])
		// add 1 for the subsequence formed by nums[i] alone
		dp[i][currParity][0] = (dp[i][currParity][0] + 1) % MOD
		for p := range 2 {
			for k := range 2 {
				// get prev count for this parity and streak
				count := dp[i-1][p][k]
				if count == 0 {
					continue
				}
				// carry forward the previous count
				dp[i][p][k] = (dp[i][p][k] + count) % MOD

				if currParity != p {
					// new streak starts in this case, so update k=0
					dp[i][currParity][0] = (dp[i][currParity][0] + count) % MOD
				} else {
					if k == 0 {
						// can only add for k=0 case, otherwise we get a 3 length streak
						dp[i][currParity][1] = (dp[i][currParity][1] + count) % MOD
					}
				}
			}
		}
	}

	ans := 0
	for p := range 2 {
		for k := range 2 {
			ans = (ans + dp[n-1][p][k]) % MOD
		}
	}
	return ans
}

func parity(val int) int {
	if val%2 == 0 {
		return 0
	}
	return 1
}
