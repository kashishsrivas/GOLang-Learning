package main

import "fmt"

func main() {
	marks := 82
	if marks >= 90 {
		fmt.Println("A")
	} else if marks >= 75 {
		fmt.Println("B")
	} else if marks >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("Fail")
	}
}
