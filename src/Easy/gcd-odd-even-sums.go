// https://leetcode.com/contest/weekly-contest-464/problems/gcd-of-odd-and-even-sums/
package easy

func gcdOfOddEvenSums(n int) int {
	var sumOdd int
	var sumEven int
	for i := 1; i <= 2*n; i++ {
		if i%2 == 0 {
			sumEven += i
		} else {
			sumOdd += i
		}
	}
	//fmt.Println(sumEven, sumOdd)
	return gcd(sumEven, sumOdd)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
