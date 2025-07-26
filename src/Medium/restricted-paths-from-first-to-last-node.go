//https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/

// you have to find the number of restricted paths from node 1 to n such that distance[i] > distance[i+1]
// where distance[i] is the shortest distance to last node from node i

package medium

import (
	"container/heap"
	"fmt"
)

type Item struct {
	node     int
	distance int
	index    int // index in the priority queue
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq *PQ) Push(x any) {
	// place the item at the end
	n := pq.Len()
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	// the heap package places the element at the end of the slice before calling this Pop function
	n := pq.Len()
	old := *pq
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (pq *PQ) Update(item *Item, distance int) {
	item.distance = distance
	heap.Fix(pq, item.index)
}

func countRestrictedPaths(n int, edges [][]int) int {
	adj := BuildAdjacencyList(n, edges)
	// run dijkstra to get the distance to last node from all the other nodes
	// build the PQ
	pq := make(PQ, n)
	heap.Init(&pq)
	// push the node `n` with distance 0
	pq.Push(&Item{
		node:     n,
		distance: 0,
		index:    0,
	})
	distances := make([]int, n)
	distances[n-1] = 0
	RunDijkstra(&distances, adj, &pq, n)
	// now the distances slice contains the shortest distance from node i to node n
	fmt.Println(distances)
	return -1
}

func RunDijkstra(distances *[]int, adj map[int][]Edge, pq *PQ, n int) {
	// pop the element with least distance until pq is empty
	// relax all its neighbours
	distArr := *distances
	for pq.Len() > 0 {
		x := heap.Pop(pq)
		curr := x.(Item)
		// go to all connected edges
		for _, edge := range adj[curr.node] {
			newDist := distArr[curr.node] + edge.w
			// if the new distance is less than the curr distance then update it
			if newDist < distArr[edge.dest] {
				distArr[edge.dest] = newDist
				//update it in PQ
				heap.Push(pq, &Item{
					node:     edge.dest,
					distance: newDist,
					index:    -1,
				})
			}
		}
	}
}

func BuildAdjacencyList(n int, edges [][]int) map[int][]Edge {
	adj := make(map[int][]Edge)
	for _, edge := range edges {
		source := edge[0]
		dest := edge[1]
		w := edge[2]
		adj[source] = append(adj[source], Edge{
			dest: dest,
			w:    w,
		})
	}
	return adj
}

type Edge struct {
	dest, w int
}
