package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	// in some cases, we need to do type conversions...
	// we need to use := unless we want to use a direct variable
	// it's better to use a special assignment operator
	// we are using it as one since they are the same type so we don't need type conversion
	// variables can be reassigned
	// does it have a garbage collector?
	// pointer use is as important as variables

	fmt.Print("Please enter your investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter your years: ")
	fmt.Scan(&years)

	fmt.Print("Please enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// investmentAmount = 2000

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Print("The first future value is: ")
	//fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.1f\nFuture Real Value: %.1f", futureValue, futureRealValue)
	// fmt.Print("The second real future value is: ")
	//fmt.Println("Future Real Value:", futureRealValue)
}
