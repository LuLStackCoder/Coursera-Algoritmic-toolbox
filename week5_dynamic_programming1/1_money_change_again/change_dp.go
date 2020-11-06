package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getChange(n int) int {
	money := make([]int, n)
	coins := []int{1, 3, 4}

	for i := 1; i < n; i++ {
		money[i] = int(^uint(0) >> 1)
		for _, v := range coins {
			if i >= v {
				numChange := money[i-v] + 1
				money[i] = min(numChange, money[i])
			}
		}
	}

	return money[n-1]
}

func main() {
	var n int

	fmt.Scanln(&n)

	fmt.Println(getChange(n + 1))
}
