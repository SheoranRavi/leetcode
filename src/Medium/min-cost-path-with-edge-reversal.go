// https://leetcode.com/contest/biweekly-contest-163/problems/minimum-cost-path-with-edge-reversals/

package medium

import "container/heap"

type Edge163 struct {
	to   int
	cost int
}

type Node struct {
	distance int64
	idx      int
}

type PQ163 []Node //suffixing the contest number to avoid redeclaration errors

func (pq PQ163) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PQ163) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ163) Len() int {
	return len(pq)
}

func (pq *PQ163) Push(x any) {
	*pq = append(*pq, x.(Node))
}

func (pq *PQ163) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func minCost(n int, edges [][]int) int {
	// reverse all edges to make a bi-directional graph
	// compute shortest path
	// this works because in shortest path you would not traverse the same node or edge more than once.
	const inf int64 = 1 << 60
	adj := make([][]Edge163, n)
	for _, edge := range edges {
		s, d, w := edge[0], edge[1], edge[2]
		adj[s] = append(adj[s], Edge163{to: d, cost: w})
		adj[d] = append(adj[d], Edge163{to: s, cost: 2 * w})
	}

	distance := make([]int64, n)
	for i := range distance {
		distance[i] = inf
	}
	distance[0] = 0
	pq := &PQ163{}
	heap.Init(pq)
	heap.Push(pq, Node{distance: 0, idx: 0})
	for pq.Len() > 0 {
		x := heap.Pop(pq)
		node := x.(Node)
		if node.idx == n-1 {
			return int(node.distance)
		}
		for _, edge := range adj[node.idx] {
			dist := node.distance + int64(edge.cost)
			if dist < distance[edge.to] {
				heap.Push(pq, Node{distance: dist, idx: edge.to})
				distance[edge.to] = dist
			}
		}
	}
	return -1
}
