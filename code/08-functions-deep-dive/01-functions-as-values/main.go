package main

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double) // Using a pointer to the slice allows us to modify the original slice if needed.
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers) // This function returns a transformation function based on the first element of the slice.
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1) // Using the transformation function returned by getTransformerFunction.
	moreTransformedNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

// transformNumbers takes a pointer to a slice of integers and a transformation function.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	// Iterate over the slice and apply the transformation function to each element.
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// getTransformerFunction returns a transformation function based on the first element of the slice.
func getTransformerFunction(numbers *[]int) transformFn {
	// If the first element is 1, return the double function; otherwise, return the triple function.
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// double and triple are transformation functions that can be passed to transformNumbers.
func double(number int) int {
	return number * 2
}

// triple is another transformation function.
func triple(number int) int {
	return number * 3
}
