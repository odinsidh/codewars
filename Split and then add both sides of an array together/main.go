package main

import (
	"fmt"
)

var console = fmt.Println
var consolef = fmt.Printf

func SplitAndAdd(input []int, limit int) []int {
	return coreLogick(input, limit, 0)
}

func coreLogick(input []int, limit int, level int) []int {
	if level >= limit || len(input) == 1 {
		return input
	}
	var (
		cut   int = len(input) / 2
		left  []int
		right []int
	)
	left = input[:cut]
	right = input[cut:]
	if len(input)%2 != 0 {
		left = append([]int{0}, left...)
	}
	for i := 0; i < len(right); i++ {
		left[i] = left[i] + right[i]
	}
	input = left
	return coreLogick(input, limit, level+1)
}

func main() {
	d1, d2 := []int{23, 345, 345, 345, 34536, 567, 568, 6, 34536, 54, 7546, 456}, 20
	console(SplitAndAdd(d1, d2))
}
