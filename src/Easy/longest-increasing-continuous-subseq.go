// https://leetcode.com/problems/longest-continuous-increasing-subsequence/

func findLengthOfLCIS(nums []int) int {
	// l := 0
	// r := 1
	// maxLen := 1
	// prev := nums[l]
	// for r < len(nums) {
	//     if nums[r] <= nums[l] || nums[r] <= prev {
	//         l = r
	//         r = l + 1
	//         prev = nums[l]
	//     } else{
	//         maxLen = max(maxLen, r-l+1)
	//         prev = nums[r]
	//         r++
	//     }
	// }
	// return maxLen
	// two pointer approach is not needed here, because we break entirely when the sequence starts to decrease (or equalize)
	maxLen := 1
	currLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currLen++
			maxLen = max(maxLen, currLen)
		} else {
			currLen = 1
		}
	}
	return maxLen
}