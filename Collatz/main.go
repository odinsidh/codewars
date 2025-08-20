package main

import (
	"fmt"
	"strconv"
)

var console = fmt.Println
var consolef = fmt.Printf

func Collatz(input int) (output string) {
	switch {
	case input == 1:
		return strconv.Itoa(input)
	case input%2 == 0:
		output = strconv.Itoa(input) + "->"
		input = input / 2
	case input%2 != 0:
		output = strconv.Itoa(input) + "->"
		input = input*3 + 1
	}
	return output + Collatz(input)
}

func main() {
	console(Collatz(3))
}
