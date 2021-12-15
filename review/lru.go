package review

/**
	双向链表
	map存储可寻找
	双向链表的插入与删除
 */

type Node struct {
	key int
	val int
	next *Node
	pre *Node
}

type LRUCache struct {
	head *Node
	tail *Node
	hmap map[int]*Node
	capacity int
}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		head: head,
		tail: tail,
		capacity: capacity,
		hmap: make(map[int]*Node),
	}
}


func (this *LRUCache) Get(key int) int {
	if _,ok := this.hmap[key]; !ok {
		return -1
	}

	p,_ := this.hmap[key]
	remove(p)
	insert(this.head, p)
	return p.val
}


func (this *LRUCache) Put(key int, value int)  {
	if _,ok := this.hmap[key]; !ok {
		newNode := &Node{
			key: key,
			val: value,
		}
		this.hmap[key] = newNode
		insert(this.head, newNode)
		if len(this.hmap) > this.capacity {
			delete(this.hmap, this.tail.pre.key)
			remove(this.tail.pre)
		}
	} else {
		p,_ := this.hmap[key]
		p.val = value
		remove(p)
		insert(this.head, p)
	}
}

func remove(p *Node) {
	p.pre.next = p.next
	p.next.pre = p.pre
}

func insert(p *Node, n *Node)  {
	n.next = p.next
	p.next = n
	n.pre = p
	n.next.pre = n
}

