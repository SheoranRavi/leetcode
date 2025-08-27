// https://leetcode.com/contest/weekly-contest-464/problems/partition-array-into-k-distinct-groups/description/
package medium

func partitionArray(nums []int, k int) bool {
	// n := len(nums)
	// if n % k != 0 {
	//     return false
	// }
	// return true
	// brute force
	// pick k distinct elements, put em in a group
	// visited[n] array
	// for each group create a map to see if a particular element appears more than once.
	n := len(nums)
	if n%k != 0 {
		return false
	}
	// the number of duplicates in each duplicate group must be <= the number of groups being formed
	elemsCount := make(map[int]int)
	numGroups := n / k
	for _, val := range nums {
		elemsCount[val]++
		if elemsCount[val] > numGroups {
			return false
		}
	}
	return true
}
