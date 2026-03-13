package medium

import "math"

// https://leetcode.com/problems/minimum-substring-partition-of-equal-character-frequency/description/
func minimumSubstringsInPartition(s string) int {
	// fabccddg
	// f: 1
	// a: 1
	// c: 2
	// d: 2
	// g: 1

	// fa, bc, cd, dg -> 4
	// fabc, cd,

	// abababaccddb
	// ababab
	// abab
	// abaccddb
	// 0,3;  0,5; 4,n-1
	//
	// a: 4
	// b: 4
	// c: 2
	// d: 2
	// count := make(map[byte]int)
	// return countBalancedSubstr(s, count, 0)

	// dp approach
	dp := make([]int, len(s)+1)
	for i := range len(s) + 1 {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	// dp[i] : min number of partitions for s[0..i-1]

	for i := 1; i <= len(s); i++ {
		freqMap := make(map[byte]int)
		maxFreq := 0
		for j := i; j >= 1; j-- {
			freqMap[s[j-1]]++
			maxFreq = max(maxFreq, freqMap[s[j-1]])
			numUniq := len(freqMap)
			//if maxFreq*numUniq == i-j{
			if dp[j-1] != math.MaxInt && maxFreq*numUniq == i-j+1 {
				dp[i] = min(dp[i], dp[j-1]+1)
			}
		}
		// freqMap[s[i]]--
		// if freqMap[s[i]] == 0 {
		//     delete(freqMap, s[i])
		// }
	}
	return dp[len(s)]
}

func isBalanced(freqMap map[byte]int) bool {
	if len(freqMap) == 0 {
		return true
	}
	prevV := -1
	for _, v := range freqMap {
		if prevV == -1 {
			prevV = v
		} else if v != prevV {
			return false
		}
	}
	return true
}

// minSubstr count for string starting at index i
// but if extending the currString, then this func means something else
func countBalancedSubstr(s string, count map[byte]int, i int) int {
	if i >= len(s) {
		// number of balanced substr from index i to end when i >= len(s) = 0
		return 0
	}
	count[s[i]]++
	c := count[s[i]]
	isBalanced := true
	for _, v := range count {
		if c != v {
			isBalanced = false
			break
		}
	}

	nextBalanced := math.MaxInt
	if i == len(s)-1 {
		if isBalanced {
			return 1
		}
		return math.MaxInt
	}
	if isBalanced {
		// start a new substr, only possible in this case
		newCount := make(map[byte]int)
		nextBalanced = 1 + countBalancedSubstr(s, newCount, i+1)
	}
	// or keep going with current one (possible for both cases)
	currBalanced := countBalancedSubstr(s, count, i+1)
	//fmt.Println(currBalanced, nextBalanced)
	return min(nextBalanced, currBalanced)
}
