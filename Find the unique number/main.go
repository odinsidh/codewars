package main

import (
	"fmt"
)

var console = fmt.Println
var consolef = fmt.Printf

func FindUniq(input []float32) float32 {
	var container map[float32]int = make(map[float32]int)
	var output float32
	for _, value := range input {
		container[value] += 1
	}
	for key, value := range container {
		if value == 1 {
			output = key
			break
		}
	}
	return output
}

func main() {
	var d1 []float32 = []float32{1.0, 1.0, 1.0, 2, 1.0, 1.0}
	console(FindUniq(d1))
}
