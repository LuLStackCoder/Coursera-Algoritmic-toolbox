package main

import "fmt"

func merge(a []int, b []int) ([]int, int) {
	counter := 0
	ln := len(a) + len(b)
	merged := make([]int, ln)
	i, j, k := 0, 0, 0

	for ; i < len(a) && j < len(b); k++ {
		if a[i] <= b[j] {
			merged[k] = a[i]
			i++
		} else {
			counter += len(a) - i
			merged[k] = b[j]
			j++
		}
	}

	for i < len(a) {
		merged[k] = a[i]
		i++
		k++
	}

	for j < len(b) {
		merged[k] = b[j]
		j++
		k++
	}

	return merged, counter
}

func getNumberOfInversions(a []int) (merged []int, num int) {
	if len(a) <= 1 {
		return a, num
	}

	mid := len(a) / 2

	l, lNum := getNumberOfInversions(a[:mid])
	num += lNum
	r, rNum := getNumberOfInversions(a[mid:])
	num += rNum
	merged, count := merge(l, r)
	num += count

	return
}

func main() {
	var n int

	fmt.Scanln(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	_, num := getNumberOfInversions(a)
	fmt.Println(num)
}
