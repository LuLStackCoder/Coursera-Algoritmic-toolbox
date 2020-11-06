package main

import "fmt"

func sum(a []int) (s int) {
	for i := 0; i < len(a); i++ {
		s += a[i]
	}

	return
}

func partition3(v []int) (wa int) {
	if len(v) < 3 {
		return 0
	}

	s := sum(v)
	if s%3 != 0 {
		return 0
	}

	w := s / 3

	values := make([][]int, len(v)+1)
	for i := 0; i < len(v)+1; i++ {
		values[i] = make([]int, w+1)
	}
	count := 0

	for i := 1; i < len(v)+1; i++ {
		for j := 1; j < w+1; j++ {
			values[i][j] = values[i-1][j]
			if v[i-1] <= j {
				value := values[i-1][j-v[i-1]] + v[i-1]
				if value > values[i][j] {
					values[i][j] = value
				}
			}
			if values[i][j] == w {
				count++
			}
		}
	}

	if count < 3 {
		return 0
	}

	return 1
}

func main() {
	var n int
	fmt.Scanln(&n)

	v := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println(partition3(v))
}
