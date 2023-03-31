package main

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
