// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
package hard

import "fmt"

func findSubstring(s string, words []string) []int {
	// create a map containing freq of all words
	// set window size to size of word
	// move the window over s
	// if windowString in set, then one check starts
	//  during the check put the current word into another map
	//  also count the number of matched words
	//  also check in the 2nd set if the current word is present and freq > word freq
	//      if present then start the check again
	//          re-initialize the 2nd set
	//          reset the count
	//  if count == len(words) then add the startIndex to result set
	winSize := len(words[0])
	set := make(map[string]int)
	for _, w := range words {
		set[w]++
	}
	winR := make([]rune, 0)
	for i, r := range s {
		if i == winSize-1 {
			// break one step before
			break
		}
		winR = append(winR, r)
	}
	win := string(winR)

	checking := false
	count := len(words)
	otherSet := make(map[string]int)
	res := []int{}
	startIdx := 0
	for i := winSize - 1; i < len(s); {
		if len(winR) == winSize {
			winR = winR[1:]
		}
		r := rune(s[i])
		winR = append(winR, r)
		win = string(winR)
		if freq, ok := set[win]; ok {
			if checking == false {
				startIdx = i - winSize + 1
				checking = true
			}
			if freqTrack := otherSet[win]; freqTrack == freq {
				// means it has appeared more times now
				otherSet = make(map[string]int)
				otherSet[win] = 1
				count = len(words) - 1
				i -= (freq - 1) * winSize
				startIdx = i - winSize + 1
			} else {
				count--
				otherSet[win]++
			}
			// i needs to be incremented by winSize
			// winSize - 1 runes to be slid
			for j := range winSize - 1 {
				if i+j+1 >= len(s) {
					break
				}
				winR = winR[1:]
				r = rune(s[i+j+1])
				winR = append(winR, r)
				win = string(winR)
			}
			i += winSize
		} else {
			checking = false
			i++
		}
		fmt.Println("win:", win)
		fmt.Println("curr count:", count)
		fmt.Println("other set:", otherSet)
		fmt.Println("start Idx:", startIdx)
		if count == 0 {
			res = append(res, startIdx)
			otherSet = make(map[string]int)
			count = len(words)
			// now it must start from the previous place
			startIdx += winSize
			i = startIdx
			winR = []rune{}
			for j := range winSize - 1 {
				if i+j+1 >= len(s) {
					break
				}
				r = rune(s[i+j+1])
				winR = append(winR, r)
			}
			win = string(winR)
		}
	}
	return res
}
