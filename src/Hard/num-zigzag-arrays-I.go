// https://leetcode.com/contest/weekly-contest-469/problems/number-of-zigzag-arrays-i/

package hard

func zigZagArrays(n int, l int, r int) int {
	const MOD int64 = 1000000007
	m := r - l + 1
	if n == 1 {
		return m % int(MOD)
	}
	dpUp := make([]int64, m)
	dpDown := make([]int64, m)
	// dpUp[val]: no. of arrays ending at val that have the last step as up
	// dpDown[val]: no. of arrays ending at val that have the last step as down

	// now build it up for n = 2
	for val := range m {
		dpUp[val] = int64(val) // eg, when val = 1, then 0,1 is the only option, so number of prev < val = dpUp[val]
		dpDown[val] = int64(m - 1 - val)
	}

	for i := 3; i <= n; i++ {
		// calculate new dpUp and new dpDown for length i now
		newUp := make([]int64, m)
		newDown := make([]int64, m)

		// accumulate the sum of all val endings
		prefUp := make([]int64, m)
		prefDown := make([]int64, m)
		prefUp[0] = dpUp[0]
		prefDown[0] = dpDown[0]
		for val := 1; val < m; val++ {
			prefUp[val] = (prefUp[val-1] + dpUp[val]) % MOD
			prefDown[val] = (prefDown[val-1] + dpDown[val]) % MOD
		}

		totalUp := prefUp[m-1]
		// now calculate the current dp[val]'s
		for val := 0; val < m; val++ {
			// to end with an UP into `val`, the previous step must be a DOWN into something smaller
			if val > 0 {
				newUp[val] = prefDown[val-1]
			}

			// similarly, to end with a DOWN into `val`, prev step must be UP into something larger
			// hence we calculate the number of up steps into values larger than val
			sumFromLarger := totalUp - prefUp[val]
			if sumFromLarger < 0 {
				// because of the MOD operation it can be negative
				sumFromLarger = (sumFromLarger + MOD)
			}
			newDown[val] = sumFromLarger
		}

		dpUp = newUp
		dpDown = newDown
	}

	var ans int64
	for val := 0; val < m; val++ {
		ans = (ans + dpUp[val] + dpDown[val]) % MOD
	}
	return int(ans)
}
