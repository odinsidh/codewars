package main

import (
	"fmt"
	"strings"
)

var console = fmt.Println
var consolef = fmt.Printf

type concreteWeight struct {
	camp   string
	weight int
}

type weightSignature map[rune]concreteWeight

var weight map[rune]concreteWeight = map[rune]concreteWeight{
	'w': concreteWeight{camp: "left", weight: 4},
	'p': concreteWeight{camp: "left", weight: 3},
	'b': concreteWeight{camp: "left", weight: 2},
	's': concreteWeight{camp: "left", weight: 1},
	'm': concreteWeight{camp: "right", weight: 4},
	'q': concreteWeight{camp: "right", weight: 3},
	'd': concreteWeight{camp: "right", weight: 2},
	'z': concreteWeight{camp: "right", weight: 1},
}

func AlphabetWar(input string, weight weightSignature) string {
	return define(detonate(&input), weight)
}

func detonate(input *string) string {
	var area map[int]string = make(map[int]string)
	var output []string
	for index, value := range *input {
		if value == 42 {
			area[index-1], area[index], area[index+1] = "_", "_", "_"
		}
	}
	for index, value := range *input {
		if _, ok := area[index]; !ok {
			output = append(output, string(value))
		}
	}
	return strings.Join(output, "")
}

func define(input string, weight weightSignature) string {
	var output map[string]int = map[string]int{
		"left":  0,
		"right": 0,
	}
	for _, value := range input {
		cocnreteValue := weight[value]
		output[cocnreteValue.camp] += cocnreteValue.weight
	}
	switch {
	case output["left"] > output["right"]:
		return "Left side wins!"
	case output["right"] > output["left"]:
		return "Right side wins!"
	default:
		return "Let's fight again!"
	}
}

func main() {
	var d1 string = "*wwwwww*z*"
	console(AlphabetWar(d1, weight))
}
