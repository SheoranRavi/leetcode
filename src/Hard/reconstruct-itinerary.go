package hard

import "slices"

// https://leetcode.com/problems/reconstruct-itinerary/
func findItinerary(tickets [][]string) []string {
	// eulerian path is what we need to find
	// form the adjacency list
	// then use the Hierholzer's algorithm
	// need to store edges for a node in lexical order, we will skip this first

	adj := make(map[string][]string)
	for _, t := range tickets {
		_, ok := adj[t[0]]
		if !ok {
			adj[t[0]] = make([]string, 0)
		}
		edges := adj[t[0]]
		edges = append(edges, t[1])
		// the underlying struct might change, so assign it back
		adj[t[0]] = edges
	}
	// sort each edge list
	for _, v := range adj {
		slices.Sort(v)
	}
	// now run the algorithm
	res := make([]string, 0)
	//dfs("JFK", adj, &res)
	eulerianPath("JFK", adj, &res)
	slices.Reverse(res)
	return res
}

func eulerianPath(s string, adj map[string][]string, res *[]string) {
	st := make([]string, 0, 500)
	st = append(st, s)
	for len(st) > 0 {
		curr := st[len(st)-1]
		if len(adj[curr]) > 0 {
			e := adj[curr][0]
			// remove this edge
			adj[curr] = adj[curr][1:]
			st = append(st, e)
		} else {
			// pop curr from the stack
			st = st[:len(st)-1]
			*res = append(*res, curr)
		}
	}
}

func dfs(s string, adj map[string][]string, res *[]string) {
	for len(adj[s]) > 0 {
		e := adj[s][0]
		// remove e from adj
		adj[s] = adj[s][1:]
		// perform dfs
		dfs(e, adj, res)
	}
	// add source to result
	//fmt.Println("Adding to resArr:", s)
	*res = append(*res, s)
}
