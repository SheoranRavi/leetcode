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
func longestCycleSimple(edges []int) int {
	// since there'll be a single chain or cycle from each node
	// we can track all the nodes in a map
	// map: node -> steps (steps is the number of steps to reach there)
	// if a node shows up again in the map, cycle is detected
	// just check how many steps it took to get there again to get the cycle length

	n := len(edges)
	visited := make([]bool, n)
	maxCycleLen := -1
	for i := range n {
		curr := i
		if !visited[curr] && edges[curr] != -1 {
			//another loop here to visit all nodes that can be visited from curr
			next := edges[curr]
			steps := make(map[int]int)
			steps[curr] = 0
			currSteps := 0
			for !visited[next] {
				//fmt.Println("visiting the node:", next)
				_, ok := steps[next]
				if ok {
					// this means that this node has already come before
					//fmt.Println("cycle detected at node:", next)
					// no. of steps for curr node - the no. of steps for the cycle node(next) + 1 step to reach the next node
					currLen := currSteps - steps[next] + 1
					maxCycleLen = max(maxCycleLen, currLen)
					break
				}
				steps[next] = steps[curr] + 1
				currSteps = steps[next]
				curr = next
				next = edges[next]
				if next == -1 {
					break
				}
			}
			for k, _ := range steps {
				visited[k] = true
			}
		}
	}
	return maxCycleLen
}
