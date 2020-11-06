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

func lcs2(s, t []int) int {
	matrix := make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		matrix[i] = make([]int, len(s)+1)
	}

	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			if s[j-1] == t[i-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j])
			}
		}
	}
	// fmt.Println(matrix)

	return matrix[len(t)][len(s)]
}

func main() {
	var (
		n, m int
		s, t []int
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

	fmt.Println(lcs2(s, t))
}
