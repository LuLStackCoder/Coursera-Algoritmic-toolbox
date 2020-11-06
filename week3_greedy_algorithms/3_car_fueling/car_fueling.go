package main

import "fmt"

func computeMinRefills(m int, stops []int) int {
	var curr int
	var count int

	for curr < len(stops) {
		last := curr
		for curr < len(stops)-1 && stops[curr+1]-stops[last] <= m {

			curr++
		}

		if last == curr {
			return -1
		}
		if curr >= len(stops)-1 {
			return count
		}
		count++
	}

	return count
}

func main() {
	var d, m, n int
	var stops []int
	fmt.Scanln(&d)
	fmt.Scanln(&m)
	fmt.Scanln(&n)

	stops = make([]int, n+2)
	stops[0] = 0
	for i := 1; i < n+1; i++ {
		fmt.Scan(&stops[i])
	}
	stops[n+1] = d

	fmt.Println(computeMinRefills(m, stops))
}
