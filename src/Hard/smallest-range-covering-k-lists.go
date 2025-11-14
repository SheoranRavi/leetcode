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
	heap.Init(pq)
	currMax := -1 << 32
	for i := range k {
		item := Item{
			value: nums[i][0],
			i:     i,
			j:     0,
		}
		heap.Push(pq, &item)
		currMax = max(currMax, nums[i][0])
	}
	currItem := heap.Pop(pq).(*Item)
	start := currItem.value
	end := currMax
	minDiff := end - start
	// iterate until one of the items reaches its end
	for currItem.j < len(nums[currItem.i])-1 {
		currItem.j++
		currItem.value = nums[currItem.i][currItem.j]
		heap.Push(pq, currItem)
		currMax = max(currMax, currItem.value)
		currItem = heap.Pop(pq).(*Item)
		tempStart := currItem.value
		tempEnd := currMax
		//fmt.Println("tempStart and end:", tempStart, ":", tempEnd)
		diff := tempEnd - tempStart
		if diff < minDiff {
			minDiff = diff
			start = tempStart
			end = tempEnd
		}
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