// https://leetcode.com/problems/find-all-anagrams-in-a-string/description/

package medium

func findAnagrams(s string, p string) []int {
	// two pointers
	// keep frequency of chars in p
	// keep frequency of chars in s with window length len(p)
	// when checking equality, only check for chars in p
	//     number of chars in map of s can be more than p hence this is the important step
	if len(s) < len(p) {
		return []int{}
	}
	freqP := make(map[byte]int)
	for i := range p {
		freqP[p[i]]++
	}
	r := 0
	freqS := make(map[byte]int)
	for r < len(p) {
		freqS[s[r]]++
		r++
	}
	l := 0
	res := make([]int, 0)
	if isFreqEqual(freqS, freqP) {
		res = append(res, l)
	}
	//fmt.Println("r before loop:", r)
	for r < len(s) {
		freqS[s[r]]++
		freqS[s[l]]--
		l++
		r++
		if isFreqEqual(freqS, freqP) {
			res = append(res, l)
		}
		// fmt.Println("r end of loop:", r)
		// fmt.Println("l end of loop:", l)
	}
	return res
}

func isFreqEqual(a, b map[byte]int) bool {
	// fmt.Println("a:", a)
	// fmt.Println("b:", b)
	for k, v := range b {
		if a[k] != v {
			return false
		}
	}
	return true
}
