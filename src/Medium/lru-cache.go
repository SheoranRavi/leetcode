// https://leetcode.com/problems/lru-cache/

package medium

import "container/list"

type Element = list.Element
type List = list.List

type LRUCache struct {
	nodeMap  map[int]*Element
	q        *List
	capacity int
}

type Node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		nodeMap:  make(map[int]*Element),
		q:        list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	el, ok := this.nodeMap[key]
	if !ok {
		return -1
	}
	this.updateNode(el)
	node := el.Value.(*Node)
	return node.val
}

func (this *LRUCache) updateNode(el *Element) {
	this.q.MoveToFront(el)
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// key already exists
	el, ok := this.nodeMap[key]
	if ok {
		node := el.Value.(*Node)
		node.val = value
		this.updateNode(el)
		return
	}
	// full cache
	if len(this.nodeMap) == this.capacity {
		lastEl := this.q.Back()
		node := lastEl.Value.(*Node)
		this.q.Remove(lastEl)
		delete(this.nodeMap, node.key)
	}
	// new Node
	node := &Node{
		key: key,
		val: value,
	}
	el = this.q.PushFront(node)
	this.nodeMap[key] = el
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
