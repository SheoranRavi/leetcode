// https://leetcode.com/problems/minimum-window-substring/

package hard

import "fmt"

func minWindow(s string, t string) string {
	freqT := make(map[rune]int)
	lenT := len(t)
	for _, r := range t {
		freqT[r]++
	}

	// two pointers
	// move right pointer till all are covered
	// then move left pointer till no single character is going out
	// once a character goes out then move right character till it comes back
	// repeat these steps till you reach the end of the array
	sR := []rune(s)
	endPtr := 0
	startPtr := 0
	// track the freq of elements in s
	freqS := make(map[rune]int)
	for lenT > 0 && endPtr < len(s) {
		c := sR[endPtr]
		// track freq of t runes in s
		if _, ok := freqT[c]; ok {
			freqS[c]++
			if freqS[c] <= freqT[c] {
				// use lenT to track if all t runes have been tracked
				lenT--
			}
		}
		endPtr++
	}
	fmt.Println("endPtr after covering all:", endPtr)
	if endPtr == len(s) && lenT > 0 {
		// t doesn't exist in s
		return ""
	}
	endPtr-- // to satisfy the loop invariant
	minLen := endPtr - startPtr + 1
	fmt.Println("minLen at begining: ", minLen)
	// track the min window start and ends
	resJ := endPtr
	resI := startPtr

	// move left pointer and update freqS
	// stop once the freqS of any character goes below freqT
	// then move the right pointer till that freq comes back

	// setting len(s)+1 allows us to run the left pointer logic even when endPtr has reached
	// the end
	for endPtr < len(s)+1 {
		var goingOut rune
		for ; startPtr < len(s); startPtr++ {
			goingOut = rune(s[startPtr])
			_, ok := freqS[goingOut]
			if ok {
				fmt.Printf("reducing the freq of: %c\n", goingOut)
				freqS[goingOut]--
				if freqS[goingOut] < freqT[goingOut] {
					fmt.Println("Breaking at startPtr:", startPtr)
					break
				}
			}
		}
		// at this point the min length has decreased
		currLen := endPtr - startPtr + 1
		if currLen > 0 && currLen < minLen {
			minLen = currLen
			resI = startPtr
			resJ = min(endPtr, len(s)-1)
			fmt.Println("Updated resI, resJ:", resI, resJ)
		}
		fmt.Println("startPtr after going out:", startPtr)
		startPtr++
		// now startPtr is at the char that makes t not be covered within the window
		// so we move endPtr till this going out character comes back in
		endPtr++
		for endPtr < len(s) {
			c := rune(s[endPtr])
			if _, ok := freqT[c]; ok {
				freqS[c]++
				// if the freq is equal now then it means the going out char came back
				if freqS[c] == freqT[c] {
					break
				}
			}
			endPtr++
		}
		if endPtr >= len(s) {
			break
		}
		// ADOBECODEBANC
		fmt.Println("endPtr after coming in:", endPtr)
		// now window length has increased again so we won't update minLen here
	}

	return s[resI : resJ+1]
}
