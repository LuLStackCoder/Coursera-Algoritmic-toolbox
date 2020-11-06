package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func add1(n int) int {
	return n - 1
}

func mult2(n int) int {
	if n%2 == 0 {
		return n / 2
	}

	return -1
}

func mult3(n int) int {
	if n%3 == 0 {
		return n / 3
	}

	return -1
}

func optimalSequence(n int) (int, []int) {
	a := make([]int, n)
	op := []func(int) int{add1, mult2, mult3}
	path := make([]int, 0)

	for i := 2; i < n; i++ {
		mn := int(^uint(0) >> 1)
		for _, f := range op {
			res := f(i)
			if res != -1 {
				mn = min(mn, a[res])
			}
		}
		a[i] = mn + 1

	}
	path = append(path, n-1)
	for i := n - 1; i > 0; i-- {
		mn := int(^uint(0) >> 1)
		argmn := 0
		for _, f := range op {
			res := f(i)
			if res != -1 {
				if a[res] < mn {
					mn = a[res]
					argmn = res
				}
			}
		}
		path = append(path, argmn)
	}

	return a[n-1], path
}

func main() {
	var n int

	fmt.Scanln(&n)

	m, path := optimalSequence(n + 1)
	fmt.Println(m)
	for i := 0; i < m; i++ {
		fmt.Print(path[i])
		if i != m-1 {
			fmt.Print(" ")
		}
	}
}
