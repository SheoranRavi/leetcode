// https://leetcode.com/problems/longest-non-decreasing-subarray-from-two-arrays/description/
package medium

// standard dp problem.
// main thing is to find the dp formula (obviously)
func maxNonDecreasingLength(nums1 []int, nums2 []int) int {
	// dp[i][1] : len of longest non-decreasing subarr ending with nums1[i]
	// dp[i][2] : len of longest non-decreasing subarr ending with nums2[i]

	n := len(nums1)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 3)
	}
	dp[0][1] = 1
	dp[0][2] = 1
	maxLen := 1
	for i := 1; i < n; i++ {
		if nums1[i] >= nums1[i-1] {
			dp[i][1] = dp[i-1][1] + 1
		}
		if nums1[i] >= nums2[i-1] {
			dp[i][1] = max(dp[i][1], dp[i-1][2]+1)
		}
		if nums1[i] < nums1[i-1] && nums1[i] < nums2[i-1] {
			dp[i][1] = 1
		}

		if nums2[i] >= nums2[i-1] {
			dp[i][2] = dp[i-1][2] + 1
		}
		if nums2[i] >= nums1[i-1] {
			dp[i][2] = max(dp[i][2], dp[i-1][1]+1)
		}
		if nums2[i] < nums2[i-1] && nums2[i] < nums1[i-1] {
			dp[i][2] = 1
		}
		//fmt.Println("For i:", i, "dps:", dp[i][1], dp[i][2])
		maxLen = max(maxLen, dp[i][1])
		maxLen = max(maxLen, dp[i][2])
	}

	return maxLen
}
