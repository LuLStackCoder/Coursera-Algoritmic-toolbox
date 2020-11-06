package main

import "fmt"

func fibonacci(n int) int {
	a, b := 1, 0
	for i := 0; i < n; i++ {
		a, b = a+b, a
	}
	return b
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fibonacci(n))
}
