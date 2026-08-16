package main

import "fmt"

func main() {
	age := 25
	isStudent := true

	if age >= 60 || isStudent {
		fmt.Println("Discount Available")
	} else {
		fmt.Println("No Discount")
	}
}
