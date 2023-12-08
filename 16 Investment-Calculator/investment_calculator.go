package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5 
	var investmentAmount, years float64 = 1000, 10 
	expectedReturnRate := 5.5 // declared as float64 by default
	

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	fmt.Println("Press return to continue...")
	fmt.Scanln()
}
