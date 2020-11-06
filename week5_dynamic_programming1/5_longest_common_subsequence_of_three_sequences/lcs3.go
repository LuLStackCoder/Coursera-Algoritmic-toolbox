package main

import "fmt"

func max(nums ...int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return max
}

func lcs3(s, t, u []int) int {
	matrix := make([][][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		matrix[i] = make([][]int, len(s)+1)
		for j := 0; j < len(s)+1; j++ {
			matrix[i][j] = make([]int, len(u)+1)
		}
	}

	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			for k := 1; k < len(u)+1; k++ {
				if s[j-1] == t[i-1] && t[i-1] == u[k-1] {
					matrix[i][j][k] = matrix[i-1][j-1][k-1] + 1
				} else {
					matrix[i][j][k] = max(matrix[i-1][j][k], matrix[i][j-1][k], matrix[i][j][k-1])
				}
			}
		}
	}
	// fmt.Println(matrix)

	return matrix[len(t)][len(s)][len(u)]
}

func main() {
	var (
		n, m, k int
		s, t, u []int
	)

	fmt.Scanln(&n)
	s = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	fmt.Scanln(&m)
	t = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&t[i])
	}

	fmt.Scanln(&k)
	u = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&u[i])
	}

	fmt.Println(lcs3(s, t, u))
}
