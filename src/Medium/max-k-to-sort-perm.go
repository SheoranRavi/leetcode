// https://leetcode.com/contest/weekly-contest-462/problems/maximum-k-to-sort-a-permutation/
package medium

import "math/bits"

// difficult problem if you think in the direction of finding the number of swap operations needed
// in that case you might go into the direction of constructing a graph, ANDing all the possible pairs or something like that
// main thing is that the elements are all unique in 0...n-1 range
// so, ANDing all the elements that are not in their correct place works

func sortPermutation(nums []int) int {
	n := len(nums)
	maxVal := n - 1
	k := (1 << bits.Len(uint(maxVal))) - 1 // set all the bits we need
	sorted := true
	for i := 0; i < n; i++ {
		if nums[i] != i {
			if sorted {
				sorted = false
			}
			// AND with every out of place element
			k &= nums[i]
			k &= i
		}
	}
	if sorted {
		return 0
	}
	return k
}

// another way is to start with the first element that's out of place (at least those bits would have to be set in `k`)
// then AND it with every element that's out of place
func sortPermutation2(nums []int) int {
	k := -1
	for idx, val := range nums {
		if val != idx {
			if k == -1 {
				k = val
			} else {
				k &= val
			}
		}
	}
	if k == -1 {
		return 0
	}
	return k
}
