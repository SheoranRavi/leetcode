// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
package hard

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordCount * wordLen
	if totalLen > len(s) {
		return []int{}
	}

	target := make(map[string]int)
	for _, w := range words {
		target[w]++
	}

	res := []int{}
	for offset := 0; offset < wordLen; offset++ {
		left := offset
		curr := make(map[string]int)
		// number of matches in current window
		count := 0
		for j := offset; j+wordLen <= len(s); j += wordLen {
			word := s[j : j+wordLen]
			if _, ok := target[word]; ok {
				curr[word]++

				if curr[word] <= target[word] {
					count++
				}
				// if freq of word exceeds target, need to shrink
				for curr[word] > target[word] {
					leftWord := s[left : left+wordLen]
					curr[leftWord]--
					left += wordLen
					// I do not understand this
					if curr[leftWord] < target[leftWord] {
						count--
					}
				}

				if count == wordCount {
					res = append(res, left)
					leftWord := s[left : left+wordLen]
					// continue searching by just moving the left
					curr[leftWord]--
					left += wordLen
					count--
				}
			} else {
				// reset everything
				left = j + wordLen
				curr = make(map[string]int)
				count = 0
			}
		}
	}
	return res
}
