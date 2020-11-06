package main

import "fmt"

func min(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}

	return min
}

func editDistance(s, t string) int {
	matrix := make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		matrix[i] = make([]int, len(s)+1)
		matrix[i][0] = i
	}

	for i := 0; i < len(s)+1; i++ {
		matrix[0][i] = i
	}

	for i := 1; i < len(t)+1; i++ {
		for j := 1; j < len(s)+1; j++ {
			min1 := matrix[i-1][j] + 1
			min2 := matrix[i][j-1] + 1
			min3 := matrix[i-1][j-1] + 1
			if s[j-1] == t[i-1] {
				min3--
			}
			matrix[i][j] = min(min1, min2, min3)
		}
	}

	return matrix[len(t)][len(s)]
}

func main() {
	var (
		s, t string
	)
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	fmt.Println(editDistance(s, t))
}
