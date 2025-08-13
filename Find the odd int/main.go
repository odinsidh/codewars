package main

import "fmt"

var console = fmt.Println
var consolef = fmt.Printf

func FindOdd(input []int) (output int) {
	var container map[int]int = make(map[int]int)
	for _, value := range input {
		container[value] += 1
	}
	for key, value := range container {
		if value%2 != 0 {
			output = key
		}
	}
	return
}

func main() {
	var d1 []int = []int{10}
	console(FindOdd(d1))
}
