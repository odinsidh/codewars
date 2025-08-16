// https://www.codewars.com/kata/57d165ad95497ea150000020/train/go
package main

import (
	"fmt"
)

var console = fmt.Println
var consolef = fmt.Printf

func CheckPairs(needed int, input string) (string, bool) {
	var output string
	for i := 0; i < len(input)-1; i++ {
		for _, value := range []string{"8B", "B8"} {
			if input[i] == value[0] && input[i+1] == value[1] {
				output += string(input[i : i+2])
				i++
				break
			}
		}
	}
	return output, len(output) >= needed
}

func main() {
	var d1, d2 = 7, "8j8mBliB8gimjB8B8jlB"
	console(CheckPairs(d1, d2))
}
