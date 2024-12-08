package main

import "fmt"

func main() {
	age := 32 // regulat variable

	fmt.Println("Age:", age)

	editAgeToAdultYears(&age)
	fmt.Println("Adult Years:", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
