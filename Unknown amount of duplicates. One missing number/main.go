package main

import (
	"fmt"
	"sort"
)

var console = fmt.Println
var consolef = fmt.Printf

func FindDupsMiss(input []int) (missed int, output []int) {
	sort.Sort(sort.IntSlice(input))
	var lastAppended int
	for i := 0; i < len(input)-1; i++ {
		current := input[i]
		next := input[i+1]
		switch {
		case next-current == 2:
			missed = current + 1
		case current == next:
			if lastAppended != current {
				output = append(output, current)
				lastAppended = current
			}
		}
	}
	return
}

func main() {
	var d1 []int = []int{20, 19, 6, 6, 9, 7, 17, 16, 17, 12, 5, 6, 8, 9, 10, 14, 13, 11, 14, 15, 19}
	console(FindDupsMiss(d1))
}
