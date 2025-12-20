// https://leetcode.com/problems/lfu-cache/

package hard

import (
	"container/list"
)

type List = list.List
type Element = list.Element

type LFUCache struct {
	freqMap  map[int]*List
	nodeMap  map[int]*Element
	capacity int
	minFreq  int
}

// map : freq => LL
// map : key => node

type Node struct {
	key  int
	val  int
	freq int
}

func Constructor(capacity int) LFUCache {
	lfuCache := LFUCache{
		freqMap:  make(map[int]*List),
		nodeMap:  make(map[int]*Element),
		capacity: capacity,
		minFreq:  1,
	}
	return lfuCache
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.nodeMap[key]
	if !ok {
		return -1
	}
	// update the frequency counter for this node
	this.updateFreqCounter(el)
	node := el.Value.(*Node)
	return node.val
}

// update the frequency counter for this node
func (this *LFUCache) updateFreqCounter(el *Element) {
	// take it out of the list
	node := el.Value.(*Node)
	q := this.freqMap[node.freq]
	q.Remove(el)
	// increase minFreq if needed
	if q.Len() == 0 && this.minFreq == node.freq {
		this.minFreq++
	}
	// increase counter
	node.freq++
	// add to new list at 1 higher frequency
	newList, ok := this.freqMap[node.freq]
	if !ok {
		this.freqMap[node.freq] = list.New()
		newList = this.freqMap[node.freq]
	}
	// this returns a brand new *Element, that needs to be added to nodeMap
	newEl := newList.PushFront(node)
	this.nodeMap[node.key] = newEl // replace the old element which is nowhere now
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// if key already present
	if el, ok := this.nodeMap[key]; ok {
		node := el.Value.(*Node)
		node.val = value
		this.updateFreqCounter(el)
		return
	}
	// evict if size == capacity
	if len(this.nodeMap) == this.capacity {
		// remove from the minFreq queue
		minFreqQ := this.freqMap[this.minFreq]
		lastEl := minFreqQ.Back()
		lastNode := lastEl.Value.(*Node)
		minFreqQ.Remove(lastEl) // remove last element
		delete(this.nodeMap, lastNode.key)
	}
	// add new node to cache
	node := &Node{
		key:  key,
		val:  value,
		freq: 1,
	}
	q1, ok := this.freqMap[1] // first queue
	if !ok {
		q1 = list.New()
		this.freqMap[1] = q1 // assign it DONT FORGET
	}
	newEl := q1.PushFront(node)
	this.nodeMap[key] = newEl
	this.minFreq = 1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
