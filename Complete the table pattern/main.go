package main

import (
	"fmt"
	"strings"
)

var console = fmt.Println
var consolef = fmt.Printf

func pattern(height int, wide int, data string) string {
	var output string = build(wide, height)
	var lenght int = wide * height
	var container []string = make([]string, lenght)
	for index, value := range strings.Split(data, "") {
		container[index] = value
	}
	for _, value := range container {
		if value == "" {
			value = " "
		}
		output = strings.Replace(output, "$value$", value, 1)
	}
	return strings.TrimRight(output, "\n")
}

func build(wide int, height int) string {
	var (
		output  string
		counter int
	)
	for {
		output += builder(0, wide)
		if counter == height {
			break
		}
		output += builder(1, wide)
		counter += 1
	}
	return output
}

func builder(what, amount int) string {
	var (
		pattern []string   = []string{"+---+", "| $value$ |"}
		replace [][]string = [][]string{[]string{"++", "+"}, []string{"||", "|"}}
		output  string
	)
	output = strings.Repeat(pattern[what], amount)
	output = strings.Replace(output, replace[what][0], replace[what][1], -1)
	return output + "\n"
}

func main() {
	d1, d2, d3 := 4, 3, "Nice pattern"
	console(pattern(d1, d2, d3))
	// consolef("%#v\n", pattern(d1, d2, d3))

}
