package main

import "fmt"

func fibonacciHuge(n, m int) int {
	a, b := 1, 0
	perLen := 0
	for i := 0; i < n; i++ {
		a, b = (a+b)%m, a
		perLen++
		if a == 1 && b == 0 {
			break
		}
	}
	n %= perLen
	for i := 0; i < n; i++ {
		a, b = (a+b)%m, a
		perLen++
		if a == 1 && b == 0 {
			break
		}
	}
	return b
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	fmt.Println(fibonacciHuge(n, m))
}
