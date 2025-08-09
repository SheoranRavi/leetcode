package hard

import (
	"container/heap"
	"math"
)

// tried by finding all paths to n-1, but this is obviously exponential so TLEs
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// find minTime to reach n-1
	// if minTime > maxTime return -1

	// find all paths from 0 to n-1 with their cost and their time
	// consider all paths that have time < maxTime
	// return the min cost from them
	// The graph may contain multiple edges between two nodes => just keep the minTime edge in this case

	adj := buildAdj(edges)
	// let's try DFS to solve this.
	// DFS from root
	n := len(passingFees)
	visited := make([]bool, n)
	var currCost int
	var currTime int
	minCost := dfs_min(0, adj, visited, passingFees, &currCost, &currTime, maxTime)
	return minCost
}

func buildAdj(edges [][]int) map[int][]Edge {
	adj := make(map[int][]Edge)
	for _, edge := range edges {
		s := edge[0]
		d := edge[1]
		t := edge[2]
		_, ok := adj[s]
		edge := Edge{
			s:    s,
			d:    d,
			time: t,
		}
		if !ok {
			adj[s] = []Edge{edge}
		} else {
			adj[s] = append(adj[s], edge)
		}
		_, ok = adj[d]
		edge2 := Edge{
			s:    d,
			d:    s,
			time: t,
		}
		if !ok {
			adj[d] = []Edge{edge2}
		} else {
			adj[d] = append(adj[d], edge2)
		}
	}
	return adj
}

func dfs_min(node int, adj map[int][]Edge, visited []bool, passingFees []int, currCost, currTime *int, maxTime int) int {
	n := len(visited)
	if *currTime > maxTime {
		return -1
	}
	*currCost += passingFees[node]
	if node == n-1 {
		return *currCost
	}
	visited[node] = true
	var pathCost int = -1
	for _, e := range adj[node] {
		*currTime += e.time
		if !visited[e.d] {
			x := dfs_min(e.d, adj, visited, passingFees, currCost, currTime, maxTime)
			if x != -1 {
				if pathCost == -1 {
					pathCost = x
				} else {
					pathCost = min(pathCost, x)
				}
			}
		}
		*currTime -= e.time
	}
	*currCost -= passingFees[node]
	visited[node] = false
	return pathCost
}

type Edge struct {
	s, d int
	time int
}

func minCost2(maxTime int, edges [][]int, passingFees []int) int {
	// track the minCost for each (node, time step) state
	n := len(passingFees)
	minCost := make([][]int, n)
	for i := range n {
		minCost[i] = make([]int, maxTime+1)
		for j := range maxTime + 1 {
			minCost[i][j] = math.MaxInt32
		}
	}
	minCost[0][0] = passingFees[0]
	adj := buildAdj(edges)
	item := Item{
		node: 0,
		time: 0,
		cost: passingFees[0],
	}
	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, &item)
	for pq.Len() > 0 {
		curr := heap.Pop(pq)
		currItem := curr.(*Item)
		// first time this node gets popped, it'll have the minimum cost.
		if currItem.node == n-1 {
			return currItem.cost
		}
		for _, edge := range adj[currItem.node] {
			currTime := currItem.time + edge.time
			currCost := currItem.cost + passingFees[edge.d]
			if currTime <= maxTime && currCost < minCost[edge.d][currTime] {
				heap.Push(pq, &Item{
					node: edge.d,
					time: currTime,
					cost: currCost,
				})
			}
		}
	}
	return -1
}

type Item struct {
	node, time, cost int
	index            int // index in the PQ
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
	return pq[i].cost < pq[j].cost
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
