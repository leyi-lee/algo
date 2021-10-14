package week01

/**
设计循环双端队列
641
https://leetcode-cn.com/problems/design-circular-deque/
 */

/**
需要注意
在数组k 长度上+1 存放tail
head 如果在头部插入，正常是 head-1，可能越界 head = (head-1 + len(d)) %  len(d)
tail 如果在尾部插入，正常是tail+1 ，可能越界 tail = (tail+1) % len(d)

相反如果双端队列的话，删除队尾也需要tail-1 防止越界
 */


type MyCircularDeque struct {
	arr []int
	head int
	tail int
	length int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	arr := make([]int, k + 1)
	return MyCircularDeque{
		arr: arr,
		head: 0,
		tail: 0,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() { // 满了插入失败
		return false
	}

	this.head = (this.head - 1 + len(this.arr)) % len(this.arr)
	this.arr[this.head] = value
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr[this.tail] = value
	this.tail = (this.tail + 1) % len(this.arr)
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.arr)
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + len(this.arr)) % len(this.arr)
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head] // 就是head
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	// 取末尾应该是tail的前一个
	// this.tail - 1 越界问题
	n := (this.tail - 1 + len(this.arr)) % len(this.arr)
	return this.arr[n]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail + 1) % len(this.arr) == this.head
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */