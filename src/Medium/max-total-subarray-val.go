// https://leetcode.com/contest/weekly-contest-468/problems/maximum-total-subarray-value-i/
package medium

import "math"

func maxTotalValue(nums []int, k int) int64 {
	// choose k subarrays
	// k * (maxi - mini)
	// find the max element and the min element
	maxEl := math.MinInt
	minEl := math.MaxInt
	for _, num := range nums {
		maxEl = max(num, maxEl)
		minEl = min(num, minEl)
	}

	diff := int64(maxEl - minEl)
	res := int64(k) * (diff)
	return res
}
