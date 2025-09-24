// https://leetcode.com/contest/weekly-contest-468/problems/maximum-total-subarray-value-ii/

package hard

import "container/heap"

// ---------- Segment Tree for Max ----------
type SegTreeMax struct {
	n    int
	data []int
}

func NewSegTreeMax(arr []int) *SegTreeMax {
	n := len(arr)
	size := 1
	for size < n {
		size <<= 1
	}
	data := make([]int, 2*size)
	for i := range data {
		data[i] = -1 << 31 // very small
	}
	st := &SegTreeMax{n: size, data: data}
	// add the arr elements to leaf nodes
	for i, v := range arr {
		st.data[size+i] = v
	}
	// Build up the tree by calculating parents
	for i := size - 1; i >= 1; i-- {
		st.data[i] = max(st.data[i<<1], st.data[i<<1|1])
	}
	return st
}

func (st *SegTreeMax) Query(l, r int) int {
	// [l,r]
	// go to leaf nodes
	l += st.n
	r += st.n
	res := -1 << 31
	for l <= r {
		if l&1 == 1 {
			res = max(res, st.data[l])
			l++
		}
		if r&1 == 0 {
			res = max(res, st.data[r])
			r--
		}
		l >>= 1
		r >>= 1
	}
	return res
}

// ---------- Segment Tree for Min ----------
type SegTreeMin struct {
	n    int
	data []int
}

func NewSegTreeMin(arr []int) *SegTreeMin {
	n := len(arr)
	size := 1
	for size < n {
		size <<= 1
	}
	data := make([]int, 2*size)
	for i := range data {
		data[i] = 1<<31 - 1 // very large
	}
	st := &SegTreeMin{n: size, data: data}
	for i, v := range arr {
		st.data[size+i] = v
	}
	for i := size - 1; i >= 1; i-- {
		st.data[i] = min(st.data[i<<1], st.data[i<<1|1])
	}
	return st
}

func (st *SegTreeMin) Query(l, r int) int {
	// [l,r]
	l += st.n
	r += st.n
	res := 1<<31 - 1
	for l <= r {
		if l&1 == 1 {
			res = min(res, st.data[l])
			l++
		}
		if r&1 == 0 {
			res = min(res, st.data[r])
			r--
		}
		l >>= 1
		r >>= 1
	}
	return res
}

// ---------- Item + MaxHeap ----------
type Item struct {
	diff int
	l, r int
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].diff > h[j].diff } // max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- Core logic ----------
func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	segMax := NewSegTreeMax(nums)
	segMin := NewSegTreeMin(nums)

	getDiff := func(l, r int) int {
		return segMax.Query(l, r) - segMin.Query(l, r)
	}

	h := &MaxHeap{}
	heap.Init(h)

	// start with full array
	diff := getDiff(0, n-1)
	heap.Push(h, Item{diff, 0, n - 1})

	visited := make(map[[2]int]bool)
	visited[[2]int{0, n - 1}] = true

	var res int64
	count := 0
	for count < k && h.Len() > 0 {
		it := heap.Pop(h).(Item)
		res += int64(it.diff)
		count++

		// generate children
		if it.l+1 <= it.r {
			child := [2]int{it.l + 1, it.r}
			if !visited[child] {
				visited[child] = true
				heap.Push(h, Item{getDiff(it.l+1, it.r), it.l + 1, it.r})
			}
		}
		if it.l <= it.r-1 {
			child := [2]int{it.l, it.r - 1}
			if !visited[child] {
				visited[child] = true
				heap.Push(h, Item{getDiff(it.l, it.r-1), it.l, it.r - 1})
			}
		}
	}
	return res
}
