package main

import (
	"fmt"
)

func optimalSummands(n int) (ln int, summands []int) {
	for i := 1; i <= n; i++ {
		if n-i == 0 {
			summands = append(summands, i)
		}
		if n-i <= i {
			continue
		}
		summands = append(summands, i)
		n -= i
	}

	ln = len(summands)

	return ln, summands
}

func main() {
	var n int

	fmt.Scanln(&n)

	ln, summands := optimalSummands(n)
	fmt.Println(ln)
	for i := 0; i < len(summands); i++ {
		fmt.Print(summands[i])
		if i != len(summands)-1 {
			fmt.Print(" ")
		}
	}
}
