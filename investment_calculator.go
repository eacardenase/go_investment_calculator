package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Investement amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future Value: %.2f", futureValue)
	// formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f", futureRealValue)

	// fmt.Println(formattedFV)
	// fmt.Println(formattedFRV)

	fmt.Printf(`
Future Value: %.2f	
Future Value (adjusted for Inflation): %.2f`, futureValue, futureRealValue)
}
