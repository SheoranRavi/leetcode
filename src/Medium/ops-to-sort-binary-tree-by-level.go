// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/
package medium

import (
	"cmp"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minimumOperations(root *TreeNode) int {
	// process in level order
	// do min Operations possible at each level
	// min number of swaps in an array to sort it?
	//      x - 1 -> x is the number of elements not in their correct position
	//      form a new structure that contains the x elements and their new idx
	//      sort this one and then compare the idx with the new idx

	// do level order traversal
	q := make([]*TreeNode, 0)

	levelOrder := make([][]int, 0)
	levelOrder = append(levelOrder, []int{root.Val})
	level := 1
	q = append(q, root)
	var newQ []*TreeNode
	for len(q) > 0 {
		var currLevel []int
		newQ = make([]*TreeNode, 0)
		for _, item := range q {
			if item.Left != nil {
				currLevel = append(currLevel, item.Left.Val)
				newQ = append(newQ, item.Left)
			}
			if item.Right != nil {
				currLevel = append(currLevel, item.Right.Val)
				newQ = append(newQ, item.Right)
			}
		}
		if len(currLevel) > 0 {
			levelOrder = append(levelOrder, currLevel)
		}
		level++
		q = newQ
	}
	//fmt.Println(levelOrder)
	minOps := 0
	for _, level := range levelOrder {
		minOps += levelMinOp(level)
	}
	return minOps
}

func levelMinOp(level []int) int {
	n := len(level)
	if n == 0 {
		return 0
	}

	itemArr := make([]Item, n)
	for idx, val := range level {
		itemArr[idx] = Item{
			idx: idx,
			val: val,
		}
	}
	slices.SortFunc(itemArr, func(a, b Item) int {
		return cmp.Compare(a.val, b.val)
	})
	// for getting the min operations
	visited := make([]bool, n)
	numOps := 0
	for i, item := range itemArr {
		if visited[i] {
			continue
		}
		if item.idx != i {
			cycleLen := 1
			next := item.idx
			visited[next] = true
			for itemArr[next].idx != i {
				cycleLen++
				next = itemArr[next].idx
				visited[next] = true
			}
			numOps += cycleLen
		}
	}
	return numOps
}

type Item struct {
	idx int
	val int
}
