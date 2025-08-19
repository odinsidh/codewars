package main

import (
	"fmt"
	"strings"
)

var console = fmt.Println
var consolef = fmt.Printf

func FindTheKey(messages, secrets []string) (output string) {
	var (
		knowledge      map[string]string = make(map[string]string)
		sliceM, sliceS []string          = slice(messages), slice(secrets)
	)
	for i := 0; i < len(sliceM); i++ {
		for x := 0; x < len(sliceM[i]); x++ {
			a, b := string(sliceM[i][x]), string(sliceS[i][x])
			if a != b {
				if a > b {
					a, b = b, a
				}
				knowledge[a] = b
			}
		}
	}
	var replace []string = []string{"map", "[", "]", " ", ":"}
	output = fmt.Sprint(knowledge)
	for _, value := range replace {
		output = strings.ReplaceAll(output, value, "")
	}
	return
}

func slice(input []string) []string {
	return strings.Fields(strings.Trim(fmt.Sprint(input), "[]"))
}

func main() {
	var messages = []string{"dance on the table", "hide my beers", "scouts rocks"}
	var secrets = []string{"egncd pn thd tgbud", "hked mr bddys", "scplts ypcis"}
	// var messages = []string{"ny", "uy", "yn", "rk", "ej", "o ", "wu", "wg", "jt", "ir"}
	// var secrets = []string{"ow", "gw", "wo", "ek", "rj", "n ", "yg", "yu", "jt", "me"}
	// var messages = []string{"kba", "vaz", "sbz", "kcc", "bqz", "vlq", "xhx", "uwe", "cop", "eys"}
	// var secrets = []string{"ibg", "vgz", "sbz", "icc", "bqz", "vuq", "xhx", "lwd", "cpo", "drs"}
	console(FindTheKey(messages, secrets))
}
