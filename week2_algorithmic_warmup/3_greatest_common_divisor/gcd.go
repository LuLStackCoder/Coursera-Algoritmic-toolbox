package main

import "fmt"

func gcdNaive(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(gcdNaive(a, b))
}
