package main

// Order to import multiple package
import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64
	var err error
	revenue, err = inputText("Revenue: ")
	if err != nil {
		errorPrint(err)
	}
	expenses, err = inputText("Expenses: ")
	if err != nil {
		errorPrint(err)
	}
	taxRate, err = inputText("Tax Rate: ")
	if err != nil {
		errorPrint(err)
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)
	writeToFile(ebt, profit, ratio)

}

func inputText(text string) (value float64, err error) {
	fmt.Print(text)
	fmt.Scan(&value)
	if value <= 0 {
		err = errors.New("invalid entry! Not allowed")
		return value, err
	}
	return value, nil
}

func calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {

	ebt = revenue - expenses

	profit = ebt * (1 - taxRate/100)

	ratio = ebt / profit

	return ebt, profit, ratio
}

func outputText(text string, value float64) (str string) {
	fmt.Printf("%v%.1f\n", text, value)
	str = fmt.Sprintf("%v%.1f\n", text, value)
	return str
}

func errorPrint(err error) {
	fmt.Println("ERROR!")
	fmt.Println(err)
	fmt.Println("----------------")
	panic(err) // Panic message with crash
}

func writeToFile(ebt, profit, ratio float64) {
	ebtText := outputText("Earning before Tax: ", ebt)
	profitText := outputText("Profit: ", profit)
	ratioText := outputText("Ratio: ", ratio)
	final := ebtText + profitText + ratioText
	os.WriteFile("results.txt", []byte(final), 0644)
}
