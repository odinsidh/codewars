// https://www.codewars.com/kata/5298961d9ce954d77b0003a6/train/go
package main

import "fmt"

var console = fmt.Println
var consolef = fmt.Printf

func Range(start, stop, step int) (output []int) {
	var picker int = start
	for i := start; i < stop; i++ {
		if picker >= stop {
			break
		}
		output = append(output, picker)
		picker += step
	}
	return
}

func main() {
	var d1, d2, d3 = -32, 4, 10
	console(Range(d1, d2, d3))
}
