package main

import "fmt"

func main() {
	age := 32 // regualar variable

	agePointer := &age

	agePointer2 := &agePointer
	fmt.Println("Age:", age)
	fmt.Println("Age address:", agePointer)           // Address stored in pointer
	fmt.Println("Value behind pointer:", *agePointer) // Dereferencing Pointer
	fmt.Println(**agePointer2)
	// adultYears := getAgeYears(agePointer)             // passing pointer
	// adultYears2 := getAgeYears(&age)                  // passing pointer way 2
	getAgeYears(agePointer)
	// fmt.Println("Adult Years:", adultYears)
	// fmt.Println("Adult Years:", adultYears2)
	fmt.Println("Age:", age)
}

func getAgeYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
