// https://leetcode.com/contest/weekly-contest-470/problems/remove-k-balanced-substrings/description/

package medium

import "strings"

// O(n^2/k) simple solution
func removeSubstring(s string, k int) string {
	pattern := strings.Repeat("(", k) + strings.Repeat(")", k)
	for {
		newS := strings.ReplaceAll(s, pattern, "")
		if newS == s {
			break
		}
		s = newS
	}
	return s
}
