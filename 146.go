package main

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

type DLinkedList struct {
	Head *Node
	Tail *Node
}

func (this *DLinkedList) RemoveLast() {
	if this.Tail != nil {
		this.Remove(this.Tail)
	}
}

func (this *DLinkedList) Remove(n *Node) {
	if this.Tail == this.Head {
		this.Head = nil
		this.Tail = nil
		return
	}
	if n == this.Head {
		n.Next.Pre = nil
		this.Head = n.Next
		return
	}
	if n == this.Tail {
		n.Pre.Next = nil
		this.Tail = n.Pre
		return
	}
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
}

func (this *DLinkedList) AddFirst(n *Node) {
	if this.Head == nil {
		this.Head = n
		this.Tail = n
		n.Pre = nil
		n.Next = nil
		return
	}
	n.Next = this.Head
	this.Head.Pre = n
	this.Head = n
	n.Pre = nil

}

type LRUCache struct {
	Cap     int
	Size    int
	HashMap map[int]*Node
	Cache   *DLinkedList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:     capacity,
		HashMap: map[int]*Node{},
		Cache:   &DLinkedList{},
	}
}

func (this *LRUCache) Get(k int) int {
	if node, ok := this.HashMap[k]; ok {
		this.Cache.Remove(node)
		this.Cache.AddFirst(node)
		return node.Val
	}
	return -1

}

func (this *LRUCache) Put(k, val int) {
	if node, ok := this.HashMap[k]; ok {
		this.Cache.Remove(node)
		node.Val = val
		this.Cache.AddFirst(node)
	} else {
		n := &Node{Key: k, Val: val}
		this.HashMap[k] = n
		this.Cache.AddFirst(n)
		this.Size = this.Size + 1
		if this.Size > this.Cap {
			this.Size = this.Size - 1
			delete(this.HashMap, this.Cache.Tail.Key)
			this.Cache.RemoveLast()
		}
	}
}

func main() {
}

