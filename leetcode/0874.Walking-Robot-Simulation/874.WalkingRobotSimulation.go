package leetcode

import (
	"fmt"
	"math"
)

func robotSim(commands []int, obstacles [][]int) int {
	//方向：上下左右0123
	result, curDir, mObstacles := float64(0), 0, make(map[string]bool)
	x, y, stepX, stepY := 0, 0, []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	for _, v := range obstacles {
		mObstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}
	for _, v := range commands {
		switch v {
		case -1:
			curDir = (curDir + 1) % 4
		case -2:
			curDir = (curDir + 3) % 4
		default:
			for i := 0; i < v; i++ {
				tempX, tempY := x+stepX[curDir], y+stepY[curDir]
				//碰到障碍物，不移动
				if mObstacles[fmt.Sprintf("%d-%d", tempX, tempY)] {
					break
				}
				x, y = tempX, tempY
				result = math.Max(float64(x*x+y*y), result)
			}
		}
	}
	return int(result)
}
