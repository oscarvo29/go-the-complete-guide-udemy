package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	revenue, err := getValueFromTerminal("Revenue")
	if err != nil {
		log.Fatal(err)
		return
	}

	expenses, err := getValueFromTerminal("Expenses")
	if err != nil {
		log.Fatal(err)
		return
	}
	taxRate, err := getValueFromTerminal("Tax Rate")
	if err != nil {
		log.Fatal(err)
		return
	}

	ebt, profit, ratio := calc(revenue, expenses, taxRate)

	WriteFile(ebt, profit, ratio)
}

func getValueFromTerminal(text string) (float64, error) {
	var userInput float64
	fmt.Printf("%s: ", text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 1.0, errors.New("not a valid number")
	}

	return userInput, nil
}

func calc(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func WriteFile(ebt, profit, ratio float64) {
	fileText := fmt.Sprintf("%.2f, %.2f, %.2f", ebt, profit, ratio)

	err := os.WriteFile("Calculated-values.txt", []byte(fileText), 0o644)
	if err != nil {
		panic(err)
	}
}