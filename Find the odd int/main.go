package main

import "fmt"

var console = fmt.Println
var consolef = fmt.Printf

func FindOdd(seq []int) int {
	return 0
}

func main() {
	var d1 []int = []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	console(FindOdd(d1))
}
