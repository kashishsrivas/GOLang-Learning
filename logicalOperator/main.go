package main

import "fmt"

func main() {
	age := 20
	hasID := true

	fmt.Println(age >= 18 && hasID == true)
	fmt.Println(age >= 21 && hasID == true)
	fmt.Println(age >= 21 || hasID == true)
	fmt.Println(!(age >= 18))
}
