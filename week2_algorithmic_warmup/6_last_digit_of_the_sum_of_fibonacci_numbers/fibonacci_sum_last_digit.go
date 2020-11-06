package main

import "fmt"

func fibonacciSum(n int) int {
	a, b := 1, 0
	sum := 0
	n %= 60
	for i := 0; i < n; i++ {
		a, b = (a+b)%10, a
		sum = (sum + b) % 10
	}
	return sum
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(fibonacciSum(n))
}
