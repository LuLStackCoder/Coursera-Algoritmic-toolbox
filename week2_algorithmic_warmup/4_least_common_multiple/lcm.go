package main

import "fmt"

func lcm(a, b int) int {
	mult := a * b
	for b > 0 {
		a, b = b, a%b
	}
	return mult / a
}

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println(lcm(a, b))
}
