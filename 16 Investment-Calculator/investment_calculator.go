package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000, 10 
	expectedReturnRate := 5.5 // declared as float64 by default
	

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
	fmt.Println("Press return to continue...")
	fmt.Scanln()
}
