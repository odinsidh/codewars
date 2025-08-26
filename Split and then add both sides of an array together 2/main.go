package main

import (
	"fmt"
)

var console = fmt.Println
var consolef = fmt.Printf

func SplitAndAdd(input []int, limit int) []int {
mainloop:
	for i := 0; i < limit; i++ {
		if len(input) == 1 {
			break mainloop
		}
		if len(input)%2 != 0 {
			input = append([]int{0}, input...)
		}
		for x := 0; x < len(input)/2; x++ {
			input[x] = input[x] + input[x+(len(input)/2)]
		}
		input = input[:len(input)/2]
	}
	return input
}

func main() {
	d1, d2 := []int{23, 345, 345, 345, 34536, 567, 568, 6, 34536, 54, 7546, 456}, 20
	console(SplitAndAdd(d1, d2))
}
