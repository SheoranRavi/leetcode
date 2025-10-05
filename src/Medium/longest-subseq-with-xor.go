// https://leetcode.com/contest/weekly-contest-470/problems/longest-subsequence-with-non-zero-bitwise-xor/description/

package medium

// XOR of 0 and any non-zero integer would be non-zero
// So, if the array has even 1 integer not equal to others, entire array can be selected
// if all are equal, then if odd number of integers are there then the entire array can be selected
// otherwise len(arr)-1 number of elements can be selected
// special case is all 0s. Then answer is 0.
func longestSubsequence(nums []int) int {
	anyDistinct := false
	prev := nums[0]
	n := len(nums)
	res := prev
	for i := 1; i < n; i++ {
		if nums[i] != prev {
			anyDistinct = true
		}
		res = res ^ nums[i]
		prev = nums[i]
	}
	if res != 0 {
		return n
	}
	if !anyDistinct && prev == 0 {
		// all zero elements are there
		return 0
	}
	// just remove 1 element to make the zero XOR non-zero
	return n - 1
}
