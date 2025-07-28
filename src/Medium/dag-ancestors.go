// https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/description/

package medium

import "slices"

func getAncestors(n int, edges [][]int) [][]int {
	// reverse the direction of the edges
	// now the problem becomes finding all the children of a node
	// do dfs, and mark a node as visited once processed
	descendants := make(map[int]map[int]bool)
	adj := buildAdjList(n, edges)
	visited := make([]bool, n)
	for i := range n {
		dfs(adj, descendants, i, visited)
	}
	res := make([][]int, n)
	for k, v := range descendants {
		for kk, _ := range v {
			res[k] = append(res[k], kk)
		}
		slices.Sort(res[k])
	}
	return res
}

func removeDuplicates(arr []int) []int {
	// arr is sorted
	var res []int
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	prev := arr[0]
	res = append(res, prev)
	for i := 1; i < len(arr); i++ {
		if arr[i] == prev {
			continue
		}
		prev = arr[i]
		res = append(res, prev)
	}
	return res
}

func dfs(adj map[int][]int, descendants map[int]map[int]bool, idx int, visited []bool) {
	children := adj[idx]
	//descendants[idx] = append(descendants[idx], children...)
	// keep track of descendants in map
	descendants[idx] = make(map[int]bool)
	for _, child := range children {
		descendants[idx][child] = true
		if !visited[child] {
			dfs(adj, descendants, child, visited)
		}
		for desc, _ := range descendants[child] {
			_, ok := descendants[idx][desc]
			if !ok {
				descendants[idx][desc] = true
			}
		}
		//descendants[idx] = append(descendants[idx], descendants[child]...)
	}
	visited[idx] = true
}

func buildAdjList(n int, edges [][]int) map[int][]int {
	adj := make(map[int][]int)
	for _, edge := range edges {
		s := edge[0]
		d := edge[1]
		// we're reversing the direction of edges
		adj[d] = append(adj[d], s)
	}
	return adj
}
