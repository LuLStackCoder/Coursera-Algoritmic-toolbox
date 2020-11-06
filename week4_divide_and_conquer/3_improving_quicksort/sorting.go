package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition2(a []int, l, r int) int {
	x := a[l]
	j := l

	for i := l + 1; i < r; i++ {
		if a[i] <= x {
			j++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[l], a[j] = a[j], a[l]

	return j
}

func randomizedQuickSort(a []int, l int, r int) {
	if l >= r {
		return
	}
	k := rand.Intn(r-l) + l
	a[l], a[k] = a[k], a[l]
	m := partition2(a, l, r)
	randomizedQuickSort(a, l, m)
	randomizedQuickSort(a, m+1, r)
}

func partition3(a []int, l, r int) (int, int) {
	x := a[l]
	j := l + 1
	k := l

	for i := l; i < k+1; {
		if a[i] <= x {
			k++
			a[i], a[k] = a[k], a[i]
			if a[k] < x {
				a[j], a[k] = a[k], a[j]
				j++
			}
		}
	}
	a[l], a[j] = a[j], a[l]

	return j, k
}

func randomizedQuickSort3(a []int, l int, r int) {
	if l >= r {
		return
	}
	k := rand.Intn(r-l) + l
	a[l], a[k] = a[k], a[l]

	m1, m2 := partition3(a, l, r)
	randomizedQuickSort(a, l, m1-1)
	randomizedQuickSort(a, m2+1, r)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int

	// loop:
	// 	for {
	// 		n = rand.Intn(100)
	// 		a := make([]int, n)
	// 		b := make([]int, n)
	// 		for i := 0; i < n; i++ {
	// 			l := rand.Intn(100)
	// 			a[i] = l
	// 			b[i] = l
	// 		}
	// 		// fmt.Println(a)
	// 		randomizedQuickSort(a, 0, n)
	// 		randomizedQuickSort3(b, 0, n-1)
	// 		for i := 0; i < n; i++ {
	// 			if a[i] != b[i] {
	// 				fmt.Println("a:", a)
	// 				fmt.Println("b:", b)
	// 				break loop
	// 			}
	// 		}
	// 	}

	fmt.Scanln(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	randomizedQuickSort(a, 0, n)

	for i := 0; i < n; i++ {
		fmt.Print(a[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
}
