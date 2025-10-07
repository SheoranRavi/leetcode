// https://leetcode.com/problems/maximum-strictly-increasing-cells-in-a-matrix/
package hard

import (
	"cmp"
	"slices"
)

func maxIncreasingCells(mat [][]int) int {
	// flatten the matrix into val, r, c
	m := len(mat)
	n := len(mat[0])
	flat := []Item{}
	for i := range m {
		for j := range n {
			item := Item{
				val: mat[i][j],
				r:   i,
				c:   j,
			}
			flat = append(flat, item)
		}
	}

	// sort by vals
	slices.SortFunc(flat, func(a, b Item) int {
		return cmp.Compare(a.val, b.val)
	})

	// The longest seq ending in row number r and col number j tracked in these
	rowBest := make([]int, m)
	colBest := make([]int, n)
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}
	var prev Item
	rowMap := make(map[int]int)
	colMap := make(map[int]int)
	for i, item := range flat {
		if i == 0 {
			dp[item.r][item.c] = 1
			rowMap[item.r] = 1
			colMap[item.c] = 1
			prev = item
			continue
		}
		if item.val != prev.val {
			// update the rowBest and colBest first
			for k, v := range rowMap {
				rowBest[k] = max(v, rowBest[k])
			}
			for k, v := range colMap {
				colBest[k] = max(v, colBest[k])
			}
			// reset the rowMap and colMap
			rowMap = map[int]int{}
			colMap = map[int]int{}
		}
		dp[item.r][item.c] = 1 + max(rowBest[item.r], colBest[item.c])

		rowMap[item.r] = max(rowMap[item.r], dp[item.r][item.c])
		colMap[item.c] = max(colMap[item.c], dp[item.r][item.c])
		prev = item
	}
	//fmt.Println(dp)
	res := 0
	for i := range m {
		for j := range n {
			res = max(res, dp[i][j])
		}
	}
	return res
}

type Item struct {
	val int
	r   int
	c   int
}
