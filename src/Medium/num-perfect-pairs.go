// https://leetcode.com/contest/biweekly-contest-163/problems/number-of-perfect-pairs/

package medium

import "slices"

func perfectPairs(nums []int) int64 {
	// ordering doesn't matter
	// only the first condition matters, 2nd condition is always true
	// if signs are opposite, |a+b| will be min
	// if signs are same, |a-b| will be min
	// so, signs don't matter either
	// take absolute value of all nums
	// sort it
	// evaluate whether the current is a perfect pair
	// if it is, increase j
	// if it isn't then increase i
	n := len(nums)
	nums2 := make([]int, len(nums))
	for i, v := range nums {
		if v < 0 {
			nums2[i] = -v
		} else {
			nums2[i] = v
		}
	}
	slices.Sort(nums2)
	var res int64
	j := 0
	for i := 0; i < n; i++ {
		for j < n && nums2[j]-nums2[i] <= nums2[i] {
			j++
		}
		// all indices from i+1 up to j-1 are valid with i
		res += int64(j - i - 1)
	}
	return res
}
