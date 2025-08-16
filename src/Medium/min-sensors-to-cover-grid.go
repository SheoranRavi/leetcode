// https://leetcode.com/contest/biweekly-contest-163/problems/minimum-sensors-to-cover-grid/

package medium

// I was not able to solve it during the competition
// because I am stupid

func minSensors(n int, m int, k int) int {
    if k == 0 {
        return n*m
    }
    // put one at cheb distance from the border
    // then put another at cheb distance from this one till you reach the other end
    i := min(n-1, k)
    j := min(m-1, k)

    // increase each one independently
    iLimit := increment(n, i, k)
    jLimit := increment(m, j, k)
    numSensors := iLimit*jLimit
    return numSensors
}

func increment(n int, i int, k int) int {
    res := 1
    for i < n {
        i = i + k + 1
        if i < n{
            res++
        }
        i += k
    }
    return res
}©leetcode