package main

import (
	"fmt"
	"sort"
)

func count(seq []int, num int) (counter int) {
	if num == -1 {
		return 0
	}

	for i := 0; i < len(seq); i++ {
		if num == seq[i] {
			counter++
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func GetMajorityElementNaive(seq []int) int {
	for i := 0; i < len(seq); i++ {
		counter := 0
		for j := i; j < len(seq); j++ {
			if seq[i] == seq[j] {
				counter++
			}
		}
		if counter > len(seq)/2 {
			return 1
		}
	}

	return 0
}

func GetMajorityElementSort(seq []int) int {
	if len(seq) <= 1 {
		return 1
	}

	sort.Ints(seq)

	for i := 1; i < len(seq); i++ {
		counter := 1
		for i < len(seq) && seq[i] == seq[i-1] {
			counter++
			i++
		}
		if counter > len(seq)/2 {
			return 1
		}
	}

	return 0
}

func GetMajorityElementMap(seq []int) int {
	counters := make(map[int]int)

	for i := 0; i < len(seq); i++ {
		counters[seq[i]]++
	}

	for _, v := range counters {
		if v > len(seq)/2 {
			return 1
		}
	}

	return 0
}

func getMajorityElementRec(seq []int) int {
	if len(seq) == 1 {
		return seq[0]
	}

	mid := len(seq) / 2

	l := getMajorityElementRec(seq[0:mid])
	r := getMajorityElementRec(seq[mid:])

	if l == r {
		return l
	}

	lCounter := count(seq, l)
	rCounter := count(seq, r)

	if lCounter > mid {
		return l
	} else if rCounter > mid {
		return r
	}

	return 0
}

func GetMajorityElementDiv(seq []int) int {
	res := getMajorityElementRec(seq)
	if res == 0 {
		return 0
	}

	return 1
}

func main() {
	var (
		n   int
		seq []int
	)
	// seq = []int{2, 2, 2, 2, 3, 3}
	fmt.Scanln(&n)
	seq = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&seq[i])
	}
	fmt.Println(GetMajorityElementDiv(seq))
}
