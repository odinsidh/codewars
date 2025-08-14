package main

import (
	"fmt"
	"strconv"
	"strings"
)

var console = fmt.Println
var consolef = fmt.Printf

func Numericals(input string) string {
	var (
		seen   map[rune]int = make(map[rune]int)
		output []string
	)
	for _, value := range input {
		seen[value] += 1
		output = append(output, strconv.Itoa(seen[value]))
	}
	return strings.Join(output, "")
}

func main() {
	var d1 string = "Hello, World!"
	console(Numericals(d1))
}
