package main

import (
	"fmt"
	"strings"
)

var console = fmt.Println
var consolef = fmt.Printf

func duplicate_count(input string) (output int) {
	var knowledge map[rune]int = make(map[rune]int)
	input = strings.ToLower(input)
	for _, value := range input {
		knowledge[value]++
	}
	for _, value := range knowledge {
		if value > 1 {
			output++
		}
	}
	return
}

func main() {
	var d1 string = "abcdeaB11"
	console(duplicate_count(d1))
}
