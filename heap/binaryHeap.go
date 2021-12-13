package heap

type BinaryHeap struct {
	heap []int
	head int
	tail int
}

func Construct(len int) BinaryHeap {
	return BinaryHeap{
		heap: make([]int, len + 1),
		head: 0,
		tail: 0,
	}
}


func (b *BinaryHeap) insert(val int) bool {
	if b.head == b.tail {
		b.heap[1] = val
		b.tail = 2
		return true
	}

	// 先放入末尾
	b.heap[b.tail] = val
	b.tail++
	return true
}

func (b *BinaryHeap) maxHeap() int {
	return b.heap[1]
}

func (b *BinaryHeap) pop() {

}