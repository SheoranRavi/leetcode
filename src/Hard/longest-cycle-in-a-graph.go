// https://leetcode.com/problems/longest-cycle-in-a-graph/
package hard

// ----This is the TLE solution----
// I failed to see the special property of the graph that a node having a max of 1 outgoing edge means
// that the graph can be considered a chain that may terminate or have one cycle in it

// In this solution, we do DFS from each unvisited node and track the nodes that have come so far in it.
func longestCycle(edges []int) int {
	n := len(edges)
	visited := make([]bool, n)

	// do dfs for each vertex if it isn't already visited
	maxCycleLength := -1
	for i := range n {
		if !visited[i] {
			currLen := dfs(edges, visited, i)
			maxCycleLength = max(currLen, maxCycleLength)
		}
	}
	return maxCycleLength
}

func dfs(edges []int, visited []bool, curr int) int {
	// create a set for keeping track of all nodes that have been visited in current loop
	set := make(map[int]bool)
	n := len(edges)
	parent := make([]int, n)
	for i := range n {
		parent[i] = -1
	}
	return dfsHelper(edges, visited, parent, set, curr)
}

func dfsHelper(edges []int, visited []bool, parent []int, set map[int]bool, curr int) int {
	// the curr node has already been visited before, no need to check for cycle again
	if visited[curr] {
		return -1
	}
	_, ok := set[curr]
	// if already visited in current loop, then backtrack to find when the curr element last came
	if ok {
		node := curr
		cycleLen := 1
		for parent[node] != curr {
			node = parent[node]
			cycleLen++
		}
		return cycleLen
	}
	set[curr] = true
	neighbour := edges[curr]
	if neighbour == -1 {
		return -1
	}
	parent[neighbour] = curr
	cycleLen := dfsHelper(edges, visited, parent, set, neighbour)
	visited[curr] = true
	return cycleLen
}

// ----Simpler Solution----
