package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:] // This will print the slice containing prices from index 1 to the end of the array
	fmt.Println(featuredPrices)
	highlightedPrices := featuredPrices[:1] // This will print the slice containing prices from index 1 to the end
	fmt.Println(highlightedPrices)
}
