// https://leetcode.com/contest/weekly-contest-470/problems/remove-k-balanced-substrings/description/

package medium

import (
	"fmt"
	"strings"
)

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

// solve by tracking the length of the openeing group and closing group in a stack
func removeSubstring(s string, k int) string {
	// use a stack to track lenths
	// recurse on it to combine the top elements
	curr := rune(s[0])
	t := "open"
	if curr == ')' {
		t = "close"
	}
	currNode := &node{t, 1}
	stack := []*node{currNode}
	for i := 1; i < len(s); i++ {
		curr = rune(s[i])
		t = "open"
		if curr == ')' {
			t = "close"
		}
		if currNode != nil && currNode.nType == t {
			currNode.count++
		} else {
			currNode = &node{t, 1}
			stack = append(stack, currNode)
		}
		stack = recurse(stack, k)
		if len(stack) > 0 {
			currNode = stack[len(stack)-1]
		} else {
			currNode = nil
		}
	}

	res := []rune{}
	for _, item := range stack {
		for range item.count {
			if item.nType == "open" {
				res = append(res, '(')
			} else {
				res = append(res, ')')
			}
		}
	}
	return string(res)
}

func recurse(stack []*node, k int) []*node {
	//fmt.Println("Stack len:", len(stack))
	// check the top two nodes and combine them if possible
	if len(stack) < 2 {
		return stack
	}
	// fmt.Println("Before recursing:")
	// printStack(stack)
	n := len(stack)
	top := stack[n-1]
	prev := stack[n-2]
	numToRemove := 0
	// different but wrong direction
	if top.nType != prev.nType && top.nType == "open" {
		return stack
	}
	// different but less than k
	if top.nType != prev.nType && (top.count < k || prev.count < k) {
		return stack
	}
	// different and >= k
	if top.nType != prev.nType && top.count >= k && prev.count >= k {
		top.count -= k
		prev.count -= k
	}
	// same
	if top.nType == prev.nType {
		prev.count += top.count
		numToRemove = 1
	}

	// now the removal logic
	if prev.count == 0 && top.count == 0 {
		numToRemove = 2
	} else if prev.count == 0 {
		// prev to be removed so shift top to prev
		prev.count = top.count
		prev.nType = top.nType
		numToRemove = 1
	} else {
		numToRemove = 1
	}
	stack = stack[:n-numToRemove]
	// fmt.Println("After recursing:")
	// printStack(stack)
	return recurse(stack, k)
}

func printStack(stack []*node) {
	for _, item := range stack {
		fmt.Println(*item)
	}
}

type node struct {
	nType string
	count int
}
