package week08

type Trie struct {
	root *Node
}

type Node struct {
	count int
	child map[uint8]*Node
}


func Constructor() Trie {
	return Trie{
		root: &Node{
			count: 0,
			child: make(map[uint8]*Node),
		},
	}
}


func (this *Trie) Insert(word string)  {
	this.find(word, true ,false)
}


func (this *Trie) Search(word string) bool {
	return this.find(word, false ,false)
}


func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix, false,true)
}

func (this *Trie) find(s string, isInsert bool, isPrefix bool) bool {
	cur := this.root
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := cur.child[c];!ok {
			if !isInsert {
				return false
			}

			cur.child[c] = &Node{child: make(map[uint8]*Node)}
		}

		cur = cur.child[c]
	}

	if isInsert { // 如果是插入给最后加一
		cur.count++
	}

	if isPrefix { // 如果是前缀 到这里就说明跑过来了
		return true
	}

	if cur.count > 0 { // find 大于0就是满足的
		return true
	}

	return false
}
