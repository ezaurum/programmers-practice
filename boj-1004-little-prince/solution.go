package main

func solution(startx int, starty int, endx int, endy int, cn int, circles [][]int) int {
	count := 0
	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		if (startx-x)*(startx-x)+(starty-y)*(starty-y) < r*r {
			count++
		}
		if (endx-x)*(endx-x)+(endy-y)*(endy-y) < r*r {
			count++
		}
		if (startx-x)*(startx-x)+(starty-y)*(starty-y) < r*r && (endx-x)*(endx-x)+(endy-y)*(endy-y) < r*r {
			count -= 2
		}
	}
	return count
}
