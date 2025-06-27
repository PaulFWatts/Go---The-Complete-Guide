package main

import (
	"example.com03/price-calculator/prices"
)

// This is the main function that initializes the tax rates and processes each tax rate
func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}
