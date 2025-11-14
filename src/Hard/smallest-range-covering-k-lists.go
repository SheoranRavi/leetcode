// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

package hard

import (
	"container/heap"
	"fmt"
	"strconv"
)

func smallestRange(nums [][]int) []int {
	// keep pointer at beginning of each list
	// initial range would be min to max
	// then move the smallest pointer at each iteration

	// put all pointers in a min heap
	// pop from min heap at each point
	// move till one pointer reaches the end
	// then take the min and max from the items in the heap
	// heap item: {i, j}. i to stay fixed, j to move up by 1
	k := len(nums)
	pq := &PriorityQueue{}
	maxPq := &PriorityQueue{}
	heap.Init(pq)
	heap.Init(maxPq)
	maxMap := make(map[string]bool)
	for i := range k {
		item := Item{
			value: nums[i][0],
			i:     i,
			j:     0,
		}
		heap.Push(pq, &item)
		maxItem := Item{
			value: -nums[i][0],
			i:     i,
			j:     0,
		}
		maxMap[getKey(i, 0)] = true
		heap.Push(maxPq, &maxItem)
	}
	currItem := heap.Pop(pq).(*Item)
	currMaxItem := heap.Pop(maxPq).(*Item)
	start := currItem.value
	end := -currMaxItem.value
	minDiff := end - start
	// put it back in q, we just wanted to peek it
	heap.Push(maxPq, currMaxItem)
	// iterate until one of the items reaches its end
	for currItem.j < len(nums[currItem.i])-1 {
		// invalidate this item from the max pq
		maxMap[getKey(currItem.i, currItem.j)] = false
		currItem.j++
		currItem.value = nums[currItem.i][currItem.j]
		heap.Push(pq, currItem)
		maxItem := Item{
			value: -currItem.value,
			i:     currItem.i,
			j:     currItem.j,
		}
		maxMap[getKey(currItem.i, currItem.j)] = true
		heap.Push(maxPq, &maxItem)
		currItem = heap.Pop(pq).(*Item)
		// refresh the max item (no heap.Peek method otherwise need to use that)
		currMaxItem = heap.Pop(maxPq).(*Item)
		for maxMap[getKey(currMaxItem.i, currMaxItem.j)] != true {
			fmt.Println("To be discarded from maxPq", currMaxItem.value)
			currMaxItem = heap.Pop(maxPq).(*Item)
		}
		tempStart := currItem.value
		tempEnd := -currMaxItem.value
		fmt.Println("tempStart and end:", tempStart, ":", tempEnd)
		diff := tempEnd - tempStart
		if diff < minDiff {
			minDiff = diff
			start = tempStart
			end = tempEnd
		}
		// push the max item back
		heap.Push(maxPq, currMaxItem)
	}

	return []int{start, end}
}

func getKey(i, j int) string {
	return strconv.Itoa(i) + "-" + strconv.Itoa(j)
}

// An Item is something we manage in a priority queue.
type Item struct {
	value int
	i     int
	j     int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, j int) {
	item.value = value
	item.j = j
	heap.Fix(pq, item.index)
}
