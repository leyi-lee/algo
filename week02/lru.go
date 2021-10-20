package week02

import "fmt"

type Node struct {
	Key int
	Value int
	Pre *Node
	Next *Node
}

type LRUCache struct {
	Head *Node
	Tail *Node
	Len int
	MapNode map[int]*Node
	Capacity int
}


func Constructor(capacity int) LRUCache {
	head := new(Node)
	tail := new(Node)
	head.Next = tail
	tail.Pre = head

	return LRUCache{
		Head: head,
		Tail: tail,
		Len: 0,
		MapNode: map[int]*Node{},
		Capacity: capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	// 查询是用key
	if _, ok := this.MapNode[key];!ok {
		return -1
	}

	curNode,_ := this.MapNode[key] // 找到了更新到头部

	if this.Head.Next == this.MapNode[key] {
		return curNode.Value
	}

	//fmt.Println(curNode)
	//fmt.Println(curNode.Pre.Next)
	//fmt.Println(curNode.Next)
	//fmt.Println(curNode.Next.Pre)
	remove(curNode)


	insert(this.Head, curNode)

	return curNode.Value
}


func (this *LRUCache) Put(key int, value int)  {

	if _,ok := this.MapNode[key]; ok {
		newNode := this.MapNode[key]
		newNode.Value = value
		this.MapNode[key] = newNode
		return
	}

	if this.Len == this.Capacity {
		delNode := this.Tail.Pre
		remove(delNode)
		delete(this.MapNode, delNode.Key)
		this.Len--
	}

	newNode := &Node{
		Value: value,
		Key: key,
	}
	if _,ok := this.MapNode[key]; ok {
		newNode = this.MapNode[key]
		newNode.Value = value
	}

	insert(this.Head, newNode)
	this.MapNode[key] = newNode
	this.Len++
}

func remove(p *Node)  {
	p.Pre.Next = p.Next
	p.Next.Pre = p.Pre
}

func insert(p *Node, n *Node)  {
	p.Next.Pre = n
	n.Next = p.Next
	p.Next = n
	n.Pre = p
}

func (this *LRUCache) Print()  {
	for this.Head.Next != nil {
		fmt.Println(this.Head.Key,this.Head.Value,this.Head.Pre,this.Head.Next)
		this.Head = this.Head.Next
	}
	fmt.Println(this.MapNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */