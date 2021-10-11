package LinkedList

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LList struct {
	Head *Node
	length int
}

type Methods interface {
	Insert(i int, v interface{})
	Delete(i int)
	GetLength() int
	Search(v interface{}) int
	IsNull() bool
}

func CreateNode(v interface{}) *Node {
	return &Node{
		Data: v,
		Next: nil,
	}
}

func CreateList() *LList {
	return &LList{CreateNode(nil), 0}
}

func (l *LList) Insert(i int, v interface{}) {
	newNode := CreateNode(v)
	pre := l.Head
	for start:=0; start <= i; start ++ {
		if start == i-1 {
			newNode.Next = pre.Next
			pre.Next = newNode
			l.length++
			break
		}
		pre = pre.Next
	}
}

func (l *LList) delete(i int) {
	pre := l.Head
	for start:=0; start <= i; start ++ {
		currentNode := pre.Next
		if start == i-1 {
			pre.Next = currentNode.Next
			l.length--
			break
		}
		pre = pre.Next
	}
}

func (l *LList) GetLength() int {
	return l.length
}

func (l *LList) Search(v interface{}) int {
	if l.Head.Next == nil {
		return 0
	}

	currNode := l.Head.Next
	for start := 1;start <= l.length;start++ {
		if currNode.Data == v {
			return start
		}
		currNode = currNode.Next
	}
	
	return 0
}

func (l *LList) isNull() bool {
	if l.Head.Next == nil {
		return true
	}
	return false
}

func (l *LList) LPrint() {
	if l.length == 0 {
		fmt.Println("")
		return
	}
	node := l.Head.Next
	for {
		fmt.Println(node.Data)
		if node.Next == nil {
			fmt.Println(nil)
			break
		}
		node = node.Next
	}

}