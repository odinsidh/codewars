package main

import (
	"fmt"
	"strings"
	"unicode"
)

var console = fmt.Println
var consolef = fmt.Printf

var NATO map[string]string = map[string]string{
	"A": "Alfa", "B": "Bravo", "C": "Charlie", "D": "Delta", "E": "Echo", "F": "Foxtrot", "G": "Golf", "H": "Hotel", "I": "India", "J": "Juliett", "K": "Kilo", "L": "Lima", "M": "Mike", "N": "November", "O": "Oscar", "P": "Papa", "Q": "Quebec", "R": "Romeo", "S": "Sierra", "T": "Tango", "U": "Uniform", "V": "Victor", "W": "Whiskey", "X": "Xray", "Y": "Yankee", "Z": "Zulu",
}

func ToNato(input string) string {
	var container []string
	input = strings.ToUpper(input)
	for _, value := range input {
		if find, ok := NATO[string(value)]; ok {
			container = append(container, find)
			continue
		}
		if unicode.IsPunct(value) {
			container = append(container, string(value))
		}
	}
	return strings.Join(container, " ")
}

func main() {
	var d1 string = "If you can read!"
	console(ToNato(d1))
}
