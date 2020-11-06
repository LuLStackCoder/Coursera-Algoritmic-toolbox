package main

import (
	"fmt"
)

func optimalWeightMultiple(w int, a [][2]int) int {
	values := make([]int, w+1)

	for i := 0; i < w+1; i++ {
		for j := 0; j < len(a); j++ {
			if a[j][0] <= i {
				value := values[i-a[j][0]] + a[j][1]
				if value > values[i] {
					values[i] = value
				}
			}
		}
	}

	path := make([]int, len(a))
	for i := w; i > 0; {
		for j := 0; j < len(a); j++ {
			weight := i - a[j][0]
			if weight >= 0 {
				if values[i]-a[j][1] == values[weight] {
					path[j]++
					i = weight
				}
			}
		}
	}
	fmt.Println(path)

	return values[w]
}

func optimalWeightSingle(w int, a [][2]int) int {
	values := make([][]int, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		values[i] = make([]int, w+1)
	}

	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < w+1; j++ {
			values[i][j] = values[i-1][j]
			if a[i-1][0] <= j {
				value := values[i-1][j-a[i-1][0]] + a[i-1][1]
				if value > values[i][j] {
					values[i][j] = value
				}
			}
		}
	}

	path := make([]bool, len(a))
	for i, j := len(a), w; i > 1 || j > 1; {
		weight := j - a[i-1][0]
		if weight >= 0 {
			if values[i][j] != values[i-1][j] {
				path[i-1] = true
				j = weight
			} else {
				path[i-1] = false
			}
		}
		i--
	}

	fmt.Println(path)

	return values[len(a)][w]
}

func optimalWeight(w int, a []int) int {
	values := make([][]int, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		values[i] = make([]int, w+1)
	}

	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < w+1; j++ {
			values[i][j] = values[i-1][j]
			if a[i-1] <= j {
				value := values[i-1][j-a[i-1]] + a[i-1]
				if value > values[i][j] {
					values[i][j] = value
				}
			}
		}
	}

	return values[len(a)][w]
}

func main() {
	var w, n int
	fmt.Scanln(&w, &n)

	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&a[i][0], &a[i][1])
	}

	// for i := 0; i < n; i++ {
	// 	a[i][1] = a[i][0]
	// }

	fmt.Println(optimalWeightSingle(w, a))
}
