// https://leetcode.com/contest/weekly-contest-461/problems/trionic-array-ii/

package hard

import "math"

func maxSumTrionic(nums []int) int64 {
	// find all p, q candidates
	// compute the sum using prefixSum array
	segments := computePq(nums)
	prefixSum := computePrefixSum(nums)
	//fmt.Println(segments)
	//fmt.Println(prefixSum)
	// for each segment, maximize the sum of the trionic subarray
	var maxSum int64 = math.MinInt64
	for _, seg := range segments {
		segSum := findSegmentMaxSum(seg, nums, prefixSum)
		//fmt.Println("Segment sum for segment", seg, ":", segSum)
		maxSum = max(maxSum, segSum)
	}
	return maxSum
}

func findSegmentMaxSum(seg Segment, nums []int, prefSum []int64) int64 {
	// left of p, you will decrease idx till value > 0, or stop at 1st idx if negative, or till val starts to increase again
	// right of q, calculate the maxSum using PrefixSum array at each point till it starts to decrease again (or till array ends)
	left := seg.p - 1
	prev := seg.p - 2 // for comparison with immediate left val
	for left > 0 {
		if nums[prev] <= 0 {
			// prev value is going to decrease the sum, so leave it here.
			break
		}
		if nums[prev] >= nums[left] {
			break
		}
		prev--
		left--
	}
	right := seg.q + 1
	next := seg.q + 2
	currSum := prefSum[right]
	if left > 0 {
		currSum -= prefSum[left-1]
	}
	maxSum := currSum
	for right < len(nums)-1 {
		if nums[next] <= nums[right] {
			break
		}
		right++
		next++
		currSum += int64(nums[right])
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}

func computePrefixSum(nums []int) []int64 {
	prefSum := make([]int64, len(nums))
	var currSum int64 = 0
	for idx, val := range nums {
		currSum += int64(val)
		prefSum[idx] = currSum
	}
	return prefSum
}

func computePq(nums []int) []Segment {
	// find out all indices where the first derivative changes sign
	// basically need to compare 3 values: prev, curr, next
	idx := 0
	n := len(nums)
	prev := nums[idx]
	idx++
	var currSeg Segment
	segArr := make([]Segment, 0)
	peakCame := false
	for idx < n-1 {
		curr := nums[idx]
		next := nums[idx+1]
		if curr == prev || curr == next {
			peakCame = false
			currSeg = Segment{}
		} else if !peakCame && curr > prev && curr > next {
			peakCame = true
			currSeg.p = idx
		} else if peakCame && curr < prev && curr < next {
			peakCame = false
			currSeg.q = idx
			segArr = append(segArr, currSeg)
			currSeg = Segment{}
		}
		prev = curr
		idx++
	}
	return segArr
}

type Segment struct {
	p int
	q int
}
