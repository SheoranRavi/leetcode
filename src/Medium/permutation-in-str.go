// https://leetcode.com/problems/permutation-in-string/description/
package medium

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	// arrays can be directly compared with == operator in go
	//      because arrays are value types
	// hence initialize arrays instead of slices.
	var count1 [26]int
	var count2 [26]int

	for i := range s1 {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}
	if count1 == count2 {
		return true
	}
	n1 := len(s1)
	for i := len(s1); i < len(s2); i++ {
		count2[s2[i]-'a']++
		count2[s2[i-n1]-'a']--
		if count2 == count1 {
			return true
		}
	}

	return false
}
