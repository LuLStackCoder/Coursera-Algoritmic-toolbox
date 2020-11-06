package main

import (
	"fmt"
	"sort"
)

func maxDotProduct(profit, clicks []int) (product int) {
	sort.Ints(profit)
	sort.Ints(clicks)
	for i := 0; i < len(profit); i++ {
		product += profit[i] * clicks[i]
	}

	return product
}

func main() {
	var n int
	var profit, clicks []int

	fmt.Scanln(&n)

	profit, clicks = make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&profit[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&clicks[i])
	}

	fmt.Println(maxDotProduct(profit, clicks))
}
