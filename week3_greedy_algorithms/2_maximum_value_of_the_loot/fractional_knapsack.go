package main

import (
	"fmt"
	"sort"
)

func getOptimalValue(items [][2]float64, w float64) (optValue float64) {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0]/items[i][1] > items[j][0]/items[j][1]
	})

	for i := 0; i < len(items); i++ {
		if w == 0 {
			return
		}
		weight := func(a, b float64) float64 {
			if a < b {
				return a
			}
			return b
		}(items[i][1], w)

		optValue += weight * items[i][0] / items[i][1]
		w -= weight
	}

	return
}

func main() {
	var (
		n int
		w float64
	)

	fmt.Scanln(&n, &w)

	items := make([][2]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&items[i][0], &items[i][1])
	}
	fmt.Println(getOptimalValue(items, w))
}
