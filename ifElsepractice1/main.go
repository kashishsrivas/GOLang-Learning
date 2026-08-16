package main

import "fmt"

func main() {
	age := 21
	isAdult := false
	if age >= 18 {
		isAdult = true
	}
	fmt.Println(age)
	fmt.Println(isAdult)
}
