package review


/* 使用数组，初始化自己指向自己
   合并,找到两个要合并值的根节点 find(x) find(y)  放入数组 fa[find(x)] = find(y)
   找父节点
   如果是本身就返回  fa[x] == x 说明自己就是根节点
	否则 返回 find(fa[x])
	路径压缩 fa[x] = find[fa[x]) 变多叉树，不是线性的
	
 */

var fa []int

func Construct(n int) {
	for i := 0; i < n;i++ {
		fa[i] = i
	}
}

func find(x int) int {
	if fa[x] == x {
		return x
	}

	//return find(x)
	// 路径压缩
	fa[x] = find(fa[x])
	return fa[x]
}

func union(x int, y int) {
	fa[find(x)] = find(y)
}
