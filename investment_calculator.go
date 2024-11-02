package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 5.81
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 12.0

	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
