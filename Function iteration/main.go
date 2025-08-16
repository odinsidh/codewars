// https://www.codewars.com/kata/54b679eaac3d54e6ca0008c9/go
package main

import "fmt"

var console = fmt.Println
var consolef = fmt.Printf

func getDouble(input int) int {
	return input * 2
}

func CreateIterator(inputFunction func(int) int, inputAmount int) func(int) int {
	return func(inputValue int) int {
		for i := 0; i < inputAmount; i++ {
			inputValue = inputFunction(inputValue)
		}
		return inputValue
	}
}

func main() {
	var call = CreateIterator(getDouble, 3)
	console(call(3))
}
