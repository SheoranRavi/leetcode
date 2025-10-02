// https://leetcode.com/problems/number-of-increasing-paths-in-a-grid/description/

package hard

const MOD int = 1000000007

// just modified the longest increasing path in matrix solution to find the number of paths
// instead of the length of the longest path
func countPaths(grid [][]int) int {
	paths := calculateNumIncreasingPaths(grid)
	// now take the sum of all these and then add the number of nodes
	// to count the individual nodes
	res := 0
	for _, paths := range paths {
		res = (res + paths) % MOD
	}
	return res
}

// numIncreasing paths of node = Sum over num paths of children
// paths[node] = number of paths starting at `node`
func calculateNumIncreasingPaths(matrix [][]int) []int {
	adj, nodeVals := buildAdj(matrix)
	//fmt.Println(adj)
	//fmt.Println(nodeVals)
	m := len(matrix)
	n := len(matrix[0])
	numNodes := m * n
	node := 0
	visited := make([]bool, numNodes)
	paths := make([]int, numNodes)
	// do dfs from every node while tracking the visited nodes and distance
	for node < m*n {
		dfs(node, adj, nodeVals, paths, visited)
		node++
	}
	return paths
}

// return the path length in this subtree
func dfs(node int, adj map[int][]Node, nodeVals, paths []int, visited []bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	// do dfs on all children of node
	children := adj[node]
	numPaths := 1 // 1 for itself
	for _, child := range children {
		// no need to check child val, because it's been checked during adj list build
		dfs(child.node, adj, nodeVals, paths, visited)
		numPaths = (numPaths + paths[child.node]) % MOD
	}
	paths[node] = numPaths
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
			nodeVal := matrix[i][j]
			// up
			if i > 0 {
				upNode := node - n // number of cols get substracted
				upNodeVal := matrix[i-1][j]
				// add only valid children
				if upNodeVal > nodeVal {
					adj[node] = append(adj[node], Node{upNode, upNodeVal})
				}
			}
			// down
			if i < m-1 {
				downNode := node + n
				downNodeVal := matrix[i+1][j]
				if downNodeVal > nodeVal {
					adj[node] = append(adj[node], Node{downNode, downNodeVal})
				}
			}
			// left
			if j > 0 {
				leftNode := node - 1
				leftNodeVal := matrix[i][j-1]
				if leftNodeVal > nodeVal {
					adj[node] = append(adj[node], Node{leftNode, leftNodeVal})
				}
			}
			// right
			if j < n-1 {
				rightNode := node + 1
				rightNodeVal := matrix[i][j+1]
				if rightNodeVal > nodeVal {
					adj[node] = append(adj[node], Node{rightNode, rightNodeVal})
				}
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
