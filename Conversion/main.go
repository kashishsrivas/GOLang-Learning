package main

import "fmt"

func main() {
	age := 22
	height := float64(age)

	fmt.Println(age)
	fmt.Println(height)

	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", height)

	price := 99.99
	wholePrice := int(price)

	fmt.Println(price)
	fmt.Println(wholePrice)
}
