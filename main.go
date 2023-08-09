package main

func main() {
		
}

func findMaxValueOfEquation(points [][]int, k int) int {
	minHeap := &GzHeap{gz{points[0][0],points[0][1],points[0][1] - points[0][0]}}
	heap.Init(minHeap)
	ans := -900000005
	lenP := len(points)
	heapLen := 1
	for i:=1;i<lenP;i++ {
			x,y,v,pushV := points[i][0], points[i][1],points[i][0] + points[i][1], points[i][1] - points[i][0]
			for heapLen > 0 {
					tempDian := heap.Pop(minHeap).(gz)
					heapLen--
					// fmt.Println(x,y,tempDian, x - tempDian.X,k,heapLen)
					if x - tempDian.X > k {
							continue
					}
					tempAns :=  tempDian.V + v 
					if tempAns > ans {
							ans = tempAns
					}
					heap.Push(minHeap,tempDian)
					heapLen++
					break
			}
			heap.Push(minHeap,gz{x,y,pushV})
			heapLen++
	}
	return ans
}
// 小根堆
type gz struct {
	X,Y,V int
}
type GzHeap []gz
func (gzheap GzHeap) Len() int { return len(gzheap)}
func (gzheap GzHeap) Less(i,j int) bool { return gzheap[i].V > gzheap[j].V}
func (gzheap GzHeap) Swap(i,j int) { gzheap[i],gzheap[j] = gzheap[j],gzheap[i]}
func (gzheap *GzHeap) Push(x interface{}) { *gzheap = append(*gzheap, x.(gz))}
// 返回最后一个值，删除最后一个值
func (gzheap *GzHeap) Pop() interface{} {
	old := *gzheap
	n := len(old) - 1
	x := old[n]
	*gzheap = old[ : n]
	return x
}
// TODO 返回最后一个值
func (gzheap *GzHeap) Top() interface{} {
	old := *gzheap
	n := len(old) - 1
	x := old[n]
	return x
}
	// minHeap := &GzHeap{gz{0,0,grid[0][0]}}
	// heap.Init(minHeap)
	// tempDian := heap.Pop(minHeap).(gz)
	// heap.Push(minHeap,gz{nowX,nowY,nowV})