package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age // agePointer now holds the address of the variable age

	fmt.Println("Age:", *agePointer) // dereference the pointer to get the value it points to

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // dereference the pointer to change the value it points to

}
