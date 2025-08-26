package main

import "fmt"

var console = fmt.Println
var consolef = fmt.Printf

func Solve(input [][]int) int {
	var output int = 1
	for _, v0 := range input {
		var data map[int]struct{} = map[int]struct{}{}
		for _, v1 := range v0 {
			data[v1] = struct{}{}
		}
		output = output * len(data)
	}
	return output
}

func main() {
	var d1 [][]int = [][]int{{1, 2, 3}, {3, 4, 6, 6, 7}, {8, 9, 10, 12, 5, 6}}
	console(Solve(d1))
}
