package main

func bfs(grid [][]int) int {
	type bfsNode struct {
		way        int
		startIndex int
		x, y       int
	}
	bfsQueue := make([]bfsNode, 0, 100005)
	bfsQueue[0] = bfsNode{
		0, 0, 0, 0,
	}
	lenX := len(grid)
	lenY := len(grid[0])
	ans := make([][]int, lenX)
	for i := 0; i < lenX; i++ {
		ans[i] = make([]int, lenY)
	}
	ans[0][0] = 1
	times := 2
	lenQ := 1
	xMax, yMax := lenX-1, lenY-1
	for lenQ > 0 {
		for i := 0; i < lenQ; i++ {
			node := bfsQueue[i]
			way, startIndex, x, y := node.way, node.startIndex, node.x, node.y
			k := grid[x][y]
			xStart := x + 1
			xEnd := min(x+k, xMax)
			yStart := y + 1
			yEnd := min(y+k, yMax)
			if way == 1 {
				if startIndex > xStart {
					xStart = startIndex
				}
			} else if way == 2 {
				if startIndex > yStart {
					yStart = startIndex
				}
			}
			nextStartIndex := xEnd + 1
			for j := xStart; j <= xEnd; j++ {
				if ans[j][y] != 0 {
					continue
				}
				bfsQueue = append(bfsQueue, bfsNode{
					1, nextStartIndex, j, y,
				})
				ans[j][y] = times
			}
			nextStartIndex = yEnd + 1
			for j := yStart; j <= yEnd; j++ {
				if ans[x][j] != 0 {
					continue
				}
				bfsQueue = append(bfsQueue, bfsNode{
					1, nextStartIndex, x, j,
				})
				ans[x][j] = times
			}
		}
		if ans[xMax][yMax] != 0 {
			return ans[xMax][yMax]
		}
		bfsQueue = bfsQueue[lenQ:]
		lenQ = len(bfsQueue)
		times++
	}
	return -1
}
