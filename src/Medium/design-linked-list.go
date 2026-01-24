// https://leetcode.com/problems/design-linked-list/
package medium

type MyLinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{Count: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Count {
		return -1
	}
	i := 0
	curr := this.Head
	for i != index {
		curr = curr.Next
		i++
	}
	//fmt.Println("val at idx:", index, ":", curr.Val)
	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newHead := Node{
		Val:  val,
		Next: this.Head,
	}
	if this.Head != nil {
		this.Head.Prev = &newHead
	}
	this.Head = &newHead
	if this.Tail == nil {
		this.Tail = this.Head
	}
	this.Count++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := Node{
		Val:  val,
		Prev: this.Tail,
	}
	if this.Tail != nil {
		this.Tail.Next = &newNode
	}
	this.Tail = &newNode
	if this.Head == nil {
		this.Head = this.Tail
	}
	this.Count++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Count {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Count {
		this.AddAtTail(val)
		return
	}
	i := 0
	curr := this.Head
	for i != index {
		i++
		curr = curr.Next
	}
	newNode := Node{
		Val:  val,
		Next: curr,
		Prev: curr.Prev,
	}
	// <--curr-->
	curr.Prev.Next = &newNode
	curr.Prev = &newNode
	this.Count++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//fmt.Println("this.Count", this.Count)
	if index >= this.Count {
		return
	}
	i := 0
	curr := this.Head
	for i != index {
		i++
		curr = curr.Next
	}
	if i == 0 {
		if this.Head.Next != nil {
			this.Head.Next.Prev = nil
		}
		this.Head = this.Head.Next
	} else if i == this.Count-1 {
		//fmt.Println("Current tail being deleted:", *(this.Tail))
		this.Tail.Prev.Next = nil
		this.Tail = this.Tail.Prev
		//fmt.Println("New tail:", *(this.Tail))
	} else {
		// fmt.Println(*curr)
		// fmt.Println(*(curr.Prev))
		// fmt.Println(*(curr.Next))
		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}
	this.Count--
	if this.Count == 0 {
		this.Tail = nil
		this.Head = nil
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
