// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/

package hard

// Initially I had though to do DFS or use DP
// But, in order to get the max distance of at each node, you need to know the distances of its child nodes
// Hence, DFS works best
func longestIncreasingPath(matrix [][]int) int {
	adj, nodeVals := buildAdj(matrix)
	//fmt.Println(adj)
	//fmt.Println(nodeVals)
	m := len(matrix)
	n := len(matrix[0])
	numNodes := m * n
	node := 0
	visited := make([]bool, numNodes)
	distance := make([]int, numNodes)
	// do dfs from every node while tracking the visited nodes and distance
	for node < m*n {
		dfs(node, adj, nodeVals, distance, visited)
		node++
	}
	// now take the maximum from distance slice
	var res int
	for _, v := range distance {
		res = max(res, v)
	}
	return res + 1
}

// return the path length in this subtree
func dfs(node int, adj map[int][]Node, nodeVals, dist []int, visited []bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	// do dfs on all children of node
	children := adj[node]
	maxDist := 0
	for _, child := range children {
		if child.nodeVal > nodeVals[node] {
			dfs(child.node, adj, nodeVals, dist, visited)
			maxDist = max(maxDist, dist[child.node]+1)
		}
	}
	dist[node] = maxDist
}

func buildAdj(matrix [][]int) (map[int][]Node, []int) {
	adj := make(map[int][]Node)
	// number the nodes from 0 to n*m-1
	m := len(matrix)
	n := len(matrix[0])
	node := 0
	for x := range m * n {
		adj[x] = []Node{}
	}
	nodeVals := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// up
			if i > 0 {
				upNode := node - n // number of cols get substracted
				upNodeVal := matrix[i-1][j]
				adj[node] = append(adj[node], Node{upNode, upNodeVal})
			}
			// down
			if i < m-1 {
				downNode := node + n
				downNodeVal := matrix[i+1][j]
				adj[node] = append(adj[node], Node{downNode, downNodeVal})
			}
			// left
			if j > 0 {
				leftNode := node - 1
				leftNodeVal := matrix[i][j-1]
				adj[node] = append(adj[node], Node{leftNode, leftNodeVal})
			}
			// right
			if j < n-1 {
				rightNode := node + 1
				rightNodeVal := matrix[i][j+1]
				adj[node] = append(adj[node], Node{rightNode, rightNodeVal})
			}
			nodeVals[node] = matrix[i][j]
			node++
		}
	}
	return adj, nodeVals
}

type Node struct {
	node    int
	nodeVal int
}
