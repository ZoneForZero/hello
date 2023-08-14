package tools

// 小根堆
type gz struct {
	X, Y, V int
}
type GzHeap []gz

func (gzheap GzHeap) Len() int            { return len(gzheap) }
func (gzheap GzHeap) Less(i, j int) bool  { return gzheap[i].V < gzheap[j].V }
func (gzheap GzHeap) Swap(i, j int)       { gzheap[i], gzheap[j] = gzheap[j], gzheap[i] }
func (gzheap *GzHeap) Push(x interface{}) { *gzheap = append(*gzheap, x.(gz)) }

// 返回最后一个值，删除最后一个值
func (gzheap *GzHeap) Pop() interface{} {
	old := *gzheap
	n := len(old) - 1
	x := old[n]
	*gzheap = old[:n]
	return x
}

// TODO add the top func
// 

// minHeap := &GzHeap{gz{0,0,grid[0][0]}}
// heap.Init(minHeap)
// tempDian := heap.Pop(minHeap).(gz)
// heap.Push(minHeap,gz{nowX,nowY,nowV})

// func search(){
//         index := sort.Search(len(tt),func(z int) bool {
//             // 返回符合条件的最小index
//             return tt[z] > 1
//         })
// }

// segment tree
type node struct {
	l, r          int
	pre, suf, mid int
}
type segmentTree struct {
	nodes []node
}

var str string

func (tree *segmentTree) Init(nodeNum int) {
	(*tree).nodes = make([]node, nodeNum*4)
}
func (tree *segmentTree) Maintain(nowIndex int) {
	// fmt.Println("计算", nowIndex)
	lIndex := nowIndex * 2
	tree.nodes[nowIndex].pre = tree.nodes[lIndex].pre
	tree.nodes[nowIndex].suf = tree.nodes[lIndex+1].suf
	tree.nodes[nowIndex].mid = MaxInt(tree.nodes[lIndex].mid, tree.nodes[lIndex+1].mid, tree.nodes[nowIndex].pre, tree.nodes[nowIndex].suf)
	// 可以连接
	if str[tree.nodes[lIndex].r] == str[tree.nodes[lIndex+1].l] {
		if tree.nodes[lIndex].pre == (tree.nodes[lIndex].r - tree.nodes[lIndex].l + 1) {
			tree.nodes[nowIndex].pre += tree.nodes[lIndex+1].pre
		}
		if tree.nodes[lIndex+1].suf == (tree.nodes[lIndex+1].r - tree.nodes[lIndex+1].l + 1) {
			tree.nodes[nowIndex].suf += tree.nodes[lIndex].suf
		}
		if tree.nodes[lIndex].suf+tree.nodes[lIndex+1].pre > tree.nodes[nowIndex].mid {
			tree.nodes[nowIndex].mid = tree.nodes[lIndex].suf + tree.nodes[lIndex+1].pre
		}
	}
}

func (tree *segmentTree) Build(l, r, nowIndex int) {
	tree.nodes[nowIndex] = node{l, r, 1, 1, 1}
	if l == r {
		return
	}

	mid := (l + r) / 2
	lIndex := nowIndex * 2
	tree.Build(l, mid, lIndex)
	tree.Build(mid+1, r, lIndex+1)
	tree.Maintain(nowIndex)
}

func (tree *segmentTree) Update(nowIndex, changeIndex int) {
	if tree.nodes[nowIndex].l == tree.nodes[nowIndex].r {
		return
	}
	nextTreeIndex := nowIndex * 2
	if tree.nodes[nextTreeIndex].l <= changeIndex && tree.nodes[nextTreeIndex].r >= changeIndex {
		tree.Update(nextTreeIndex, changeIndex)
	} else {
		tree.Update(nextTreeIndex+1, changeIndex)
	}
	tree.Maintain(nowIndex)
}
