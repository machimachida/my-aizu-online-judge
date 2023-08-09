package challengebook

import (
	"container/list"
)

type point struct {
	x, y int
}

var vectors = [4]point{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}

const INF = 100000000

func shortestPath(n, m int, field [][]string) int {
	distances := make(map[point]int, n*m)

	var end point
	queue := list.New()

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == "S" {
				distances[point{x: i, y: j}] = 0
				queue.PushFront(point{x: i, y: j})
				continue
			}
			if field[i][j] == "G" {
				end = point{x: i, y: j}
				distances[end] = INF
				continue
			}

			distances[point{x: i, y: j}] = INF
		}
	}

	for queue.Len() > 0 {
		el := queue.Front()
		queue.Remove(el)
		p := el.Value.(point)

		if p == end {
			break
		}

		for _, v := range vectors {
			next := point{x: p.x + v.x, y: p.y + v.y}
			if next.x < 0 || next.x >= n || next.y < 0 || next.y >= m {
				continue
			}
			if field[next.x][next.y] == "#" {
				continue
			}
			if distances[next] != INF {
				continue
			}

			distances[next] = distances[p] + 1
			queue.PushBack(next)
		}
	}

	return distances[end]
}
