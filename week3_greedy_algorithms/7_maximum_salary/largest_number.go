package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(digits []int) (lrg string) {
	sort.Slice(digits, func(i, j int) bool {
		lenI := int(math.Log10(float64(digits[i])) + 1)
		lenJ := int(math.Log10(float64(digits[j])) + 1)
		return digits[i]*int(math.Pow10(lenJ))+digits[j] > digits[j]*int(math.Pow10(lenI))+digits[i]
	})
	strDigits := make([]string, len(digits))
	for i := 0; i < len(strDigits); i++ {
		strDigits[i] = strconv.Itoa(digits[i])
	}
	lrg = strings.Join(strDigits, "")

	return
}

func main() {
	var n int

	fmt.Scanln(&n)

	digits := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&digits[i])
	}
	fmt.Println(largestNumber(digits))
}
