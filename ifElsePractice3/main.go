package main

import "fmt"

func main() {
	age := 16
	hasID := true
	isStudent := true

	if (age >= 18 && hasID) || isStudent {
		fmt.Println("Special Discount")
	} else {
		fmt.Println("No Special Discount")
	}
}
