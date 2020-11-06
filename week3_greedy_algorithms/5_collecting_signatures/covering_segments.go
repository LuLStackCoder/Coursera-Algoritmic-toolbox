package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func contain(v int, segment [2]int) bool {
	if segment[0] <= v && v <= segment[1] {
		return true
	}

	return false
}

func optimalPoints(segments [][2]int) (ln int, points []int) {
	sort.Slice(segments, func(i int, j int) bool {
		return segments[i][1] < segments[j][1]
	})

	point := segments[0][1]
	points = append(points, point)

	for i := 1; i < len(segments); i++ {
		if !contain(point, segments[i]) {
			point = segments[i][1]
			points = append(points, point)
		}
	}

	ln = len(points)

	return ln, points
}

func main() {
	var n int

	fmt.Scanln(&n)

	segments := make([][2]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &segments[i][0], &segments[i][1])
	}

	ln, points := optimalPoints(segments)
	fmt.Println(ln)
	for i := 0; i < len(points); i++ {
		fmt.Print(points[i])
		if i != len(points)-1 {
			fmt.Print(" ")
		}
	}
}
