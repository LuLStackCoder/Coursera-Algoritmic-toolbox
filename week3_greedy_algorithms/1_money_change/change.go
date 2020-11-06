package main

import "fmt"

func getExchange(m int) (n int) {
	n = m/10 + m%10/5 + m%10%5

	return
}

func main() {
	var m int
	fmt.Scanln(&m)
	fmt.Println(getExchange(m))
}
