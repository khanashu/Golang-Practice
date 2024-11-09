package main

// Order to import multiple package
import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	// var investedAmount float64= 1000
	// var expectedReturnRate float64 = 5.5
	// var yearsHold float64 = 10

	// var futureValue  = investedAmount * math.Pow(1+expectedReturnRate/100, yearsHold)
	// fmt.Print(futureValue)

	// investedAmount, yearsHold, expectedReturnRate := 1000.0, 10.0, 5.5

	var investedAmount float64 // Go will assign default value as 0.0 but incase of float64 "0.0"
	expectedReturnRate := 5.5
	// yearsHold := 10.0
	var yearsHold float64

	outputText("Investment Amount: ")
	fmt.Scan(&investedAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&yearsHold)
	// futureValue := investedAmount * math.Pow(1+expectedReturnRate/100, yearsHold)

	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, yearsHold)

	futureValue, futureRealValue := calculateFutureValues(investedAmount, expectedReturnRate, yearsHold)
	// fmt.Println("Future Value:", futureValue) // Adds Newline due to println

	fmt.Printf("Future Value: %.0f\nFuture Real Value (adjusted Inflation): %.1f\n", futureValue, futureRealValue)
	// fmt.Println("Future Real Value (adjusted Inflation):", futureRealValue)

	// Using Sprintf
	formattedFv := fmt.Sprintf("Future Value: %.1f\n", futureValue) // Stores the output in format of string
	formattedRealFv := fmt.Sprintf("Future Real Value (adjusted Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFv, formattedRealFv)

	// Using Backtick
	// fmt.Printf(`Future Value: %.0f
	// Future Real Value (adjusted Inflation): %.1f\n`, futureValue, futureRealValue) // Backtick character "``" allows to add linebreaks in print statment

}

func outputText(text string) {
	fmt.Print(text)
}

// func calculateFutureValues(investedAmount, expectedReturnRate, yearsHold float64) (float64, float64) {
// 	fv := investedAmount * math.Pow(1+expectedReturnRate/100, yearsHold)
// 	rfv := fv / math.Pow(1+inflationRate/100, yearsHold)
// 	// return investedAmount * math.Pow(1+expectedReturnRate/100, yearsHold), futureValue / math.Pow(1+inflationRate/100, yearsHold)
// 	return fv, rfv
// }

// Fubncton already creates a vvariable for me in this case
func calculateFutureValues(investedAmount, expectedReturnRate, yearsHold float64) (fv float64, rfv float64) {
	fv = investedAmount * math.Pow(1+expectedReturnRate/100, yearsHold)
	rfv = fv / math.Pow(1+inflationRate/100, yearsHold)
	// return fv, rfv

	//  OR

	return
}
