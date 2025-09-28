// https://leetcode.com/contest/weekly-contest-469/problems/split-array-with-minimum-difference/submissions/1784918113/

package medium

import "fmt"

func splitArray(nums []int) int64 {
	// 1, 4, 6
	// 6, 5, 2

	// 1, 3, 7, 10
	// 10, 9, 7, 3

	// figure out validity
	turnPoint := -1
	incOnly := false
	totalSum := int64(nums[0])
	for i := 1; i < len(nums); i++ {
		if turnPoint == -1 && nums[i] <= nums[i-1] {
			turnPoint = i
		} else if turnPoint != -1 && nums[i] >= nums[i-1] {
			return -1
		}
		totalSum += int64(nums[i])
	}
	if turnPoint == -1 {
		// if no turnpoint found that means it's a increasing array
		incOnly = true
		turnPoint = len(nums) - 1
	}
	// only two possibilities need to be considered
	// calculate prefix sums
	n := len(nums)
	leftPref := make([]int64, n)
	rightPref := make([]int64, n)
	for i, v := range nums {
		if i != 0 {
			leftPref[i] = leftPref[i-1] + int64(v)
			rightPref[i] = rightPref[i-1] - int64(nums[i-1])
		} else {
			leftPref[i] = int64(v)
			rightPref[i] = totalSum
		}
	}

	fmt.Println("Turnpoint:", turnPoint)
	// edge case 1,2
	diff1 := leftPref[turnPoint-1] - rightPref[turnPoint]
	diff1 = AbsDiff(diff1)
	// if turnPoint is 1 then it's a strictly decreasing array
	// for incOnly we need to consider only 1 turnPoint
	// 3rd is the equal case where only 1 split is possible
	if turnPoint == 1 || incOnly || nums[turnPoint] == nums[turnPoint-1] {
		// means it's a decreasing/increasing array all the way
		// and we cannot take empty subarrays
		return diff1
	}
	turnPoint--
	diff2 := leftPref[turnPoint-1] - rightPref[turnPoint]
	diff2 = AbsDiff(diff2)

	return min(diff1, diff2)

}

func AbsDiff(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
