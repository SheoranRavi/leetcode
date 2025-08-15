// https://leetcode.com/contest/weekly-contest-462/problems/maximum-total-from-optimal-activation-order/
package medium

import (
	"cmp"
	"container/heap"
	"slices"
)

func maxTotal(value []int, limit []int) int64 {
	// sort by limit ascending, and value descending
	// process the list
	// put each in a min heap by limit
	// if heap size == heap top limit
	//   pop smallest limit from heap and add that to value
	// at end of loop, add all remaining elements to result
	n := len(value)
	pairArr := make([]Pair, n)
	for i := range n {
		pairArr[i] = Pair{
			limit: limit[i],
			value: value[i],
		}
	}
	slices.SortFunc(pairArr, sortFunc)
	//fmt.Println(pairArr)
	var res int64
	pq := &PriorityQueue{}
	heap.Init(pq)
	prevLimit := -1
	for _, val := range pairArr {
		if prevLimit != -1 && val.limit == prevLimit {
			// if prevLimit is set this means that limit has been deactivated.
			continue
		}
		heap.Push(pq, &val)
		top := pq.Peek()
		topElement := top.(*Pair)
		if pq.Len() == topElement.limit {
			//fmt.Println("popping and adding:", topElement)
			prevLimit = topElement.limit
			res += int64(topElement.value)
			heap.Pop(pq)
		} else {
			prevLimit = -1
		}
	}
	for pq.Len() > 0 {
		top := pq.Pop()
		topElement := top.(*Pair)
		res += int64(topElement.value)
	}
	return res
}

func sortFunc(a, b Pair) int {
	if a.limit == b.limit {
		return cmp.Compare(b.value, a.value)
	}
	return cmp.Compare(a.limit, b.limit)
}

type Pair struct {
	limit, value int
	index        int // index in the PQ
}
type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].limit < pq[j].limit
}

func (pq *PriorityQueue) Push(x any) {
	// place the item at the end
	n := pq.Len()
	item := x.(*Pair)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	// the heap package places the element at the end of the slice before calling this Pop function
	n := pq.Len()
	old := *pq
	item := old[n-1]
	item.index = -1
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Peek() any {
	// custom peek method
	n := pq.Len()
	if n == 0 {
		panic("heap is empty, can't peek")
	}
	old := *pq
	item := old[0]
	return item
}
