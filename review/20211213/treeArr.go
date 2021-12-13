package _0211213

import "fmt"

type TreeArr struct {
	n int
	a []int
	c []int
}

func NewTreeArr(nums []int) *TreeArr {
	n := len(nums)
	a := make([]int, n + 1)
	c := make([]int, n + 1)
	tree := &TreeArr{
		n: n,
		a: a,
		c: c,
	}
	for i := 1; i <= n; i++ {
		tree.a[i] = nums[i - 1]
		tree.add(i, a[i])
	}
	return tree
}

func (ta *TreeArr) add(pos int, x int)  {
	for ;pos <= ta.n; pos += lowbit(x) {
		fmt.Println(pos, lowbit(x))
		ta.c[pos] += x
	}
}

func (ta *TreeArr) Get(index int) int {
	return ta.a[index]
}

func (ta *TreeArr) Set(index int, a int) {
	ta.a[index] = a
}

func (ta *TreeArr) SumPrefix(pos int) int {
	ans := 0
	for ;pos > 0; pos -= lowbit(pos) {
		ans += ta.c[pos]
	}
	return ans
}

func lowbit(x int) int {
	return x & -x
}