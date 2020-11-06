package main

import "fmt"

func binarySearch(a []int, x int) int {
	left, right := 0, len(a)

	for left < right {
		mid := (left + right) / 2
		if a[mid] == x {
			return mid
		} else if a[mid] > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}

func linearSearch(a []int, x int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return i
		}
	}

	return -1
}

func main() {
	var (
		n int
		k int
	)
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)
	b := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < k; i++ {
		fmt.Print(binarySearch(a, b[i]))
		if i != k-1 {
			fmt.Print(" ")
		}
		// fmt.Print(linearSearch(a, b[i]))
		// if i != k-1 {
		// 	fmt.Print(" ")
		// }
	}
}
