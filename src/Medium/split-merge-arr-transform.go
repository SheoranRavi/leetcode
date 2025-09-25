// https://leetcode.com/problems/split-and-merge-array-transformation/
// we try to model the transformations as a graph
// each permutation of the array is a node, nums1 is the source, nums2 is the target
// each operation is an edge from one node to another
// we have to find the shortest distance

package medium

import (
	"strconv"
	"strings"
)

var best int

func minSplitMerge(nums1 []int, nums2 []int) int {
	if len(nums1) != len(nums2) {
		return -1
	}
	memo := make(map[string]int)
	best = 1 << 30
	numOps := dfs(nums1, nums2, memo, 0)
	return numOps
}

// this results in TLE
func dfs(arr, target []int, memo map[string]int, depth int) int {
	if depth >= best {
		return 1 << 30
	}
	if isEqual(arr, target) {
		if depth < best {
			best = depth
		}
		return 0
	}
	key := toStr(arr)
	if v, ok := memo[key]; ok && v <= depth {
		return 1 << 30
	}
	memo[key] = depth

	minOps := 1 << 30
	n := len(arr)

	for L := 0; L < n; L++ {
		for R := L; R < n; R++ {
			sub := append([]int(nil), arr[L:R+1]...)
			remArr := append([]int(nil), arr[:L]...)
			remArr = append(remArr, arr[R+1:]...)
			// put sub at every position in the remArr
			for i := 0; i <= len(remArr); i++ {
				newArr := append([]int(nil), remArr[:i]...)
				newArr = append(newArr, sub...)
				newArr = append(newArr, remArr[i:]...)

				ops := dfs(newArr, target, memo, depth+1)
				if ops+1 < minOps {
					minOps = ops + 1
				}
			}
		}
	}

	return minOps
}

func isEqual(arr, target []int) bool {
	if len(arr) != len(target) {
		return false
	}
	for i := range arr {
		if arr[i] != target[i] {
			return false
		}
	}
	return true
}

func toStr(arr []int) string {
	sb := strings.Builder{}
	for i, v := range arr {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}
