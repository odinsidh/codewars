package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var console = fmt.Println
var consolef = fmt.Printf

type balanceBook struct {
	rawInput string
	balance  float64
	expenses []expense
}

type expense struct {
	id       string
	category string
	cost     float64
}

func getBalanceBook(input string) *balanceBook {
	return &balanceBook{rawInput: input}
}

func (self *balanceBook) run() string {
	self.getBalance()
	self.getExpenses()
	return self.getOutput()
}

func Balance(input string) string {
	return getBalanceBook(input).run()
}

func (self *balanceBook) getBalance() {
	var (
		re *regexp.Regexp = regexp.MustCompile(`\d+.\d+`)
	)
	cast, err := strconv.ParseFloat(re.FindAllString(self.rawInput, 1)[0], 64)

	if err != nil {
		log.Fatal(err)
	}
	self.balance = cast
}

func (self *balanceBook) getExpenses() {
	var (
		re *regexp.Regexp = regexp.MustCompile(`(\d+)(?:\s+)([a-zA-Z]+)(?:.*?)(\d+.\d+)`)
	)
	for _, value := range re.FindAllStringSubmatch(self.rawInput, -1) {
		id := value[1]
		category := value[2]
		cost, err := strconv.ParseFloat(value[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		self.expenses = append(self.expenses, expense{
			id:       id,
			category: category,
			cost:     cost,
		})
	}
}

func (self *balanceBook) getOutput() string {
	var (
		output         []string
		expensesLen    int = len(self.expenses)
		expensesTotal  float64
		balanceCurrent float64 = self.balance
	)
	output = append(output, fmt.Sprintf("Original Balance: %.2f", self.balance))
	for _, concreteExpense := range self.expenses {
		expensesTotal += concreteExpense.cost
		balanceCurrent -= concreteExpense.cost
		output = append(output, fmt.Sprintf("%s %s %.2f Balance %.2f",
			concreteExpense.id,
			concreteExpense.category,
			concreteExpense.cost,
			balanceCurrent))
	}
	output = append(output, fmt.Sprintf("Total expense  %.2f", expensesTotal))
	output = append(output, fmt.Sprintf("Average expense  %.2f", expensesTotal/float64(expensesLen)))
	return strings.Join(output, "\n")
}

func main() {
	b1 := `
1233.00
125 Hardware;! 24.8?;
123 Flowers 93.5
127 Meat 120.90
120 Picture 34.00
124 Gasoline 11.00
123 Photos;! 71.4?;
122 Picture 93.5
132 Tyres;! 19.00,?;
129 Stamps 13.6
129 Fruits{} 17.6
129 Market;! 128.00?;
121 Gasoline;! 13.6?;
`
	console(Balance(b1))
}
