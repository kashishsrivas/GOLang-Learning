package main

import "fmt"

func main() {
	marks := 32
	var passed bool

	if marks >= 40 {
		passed = true
	} else {
		passed = false
	}

	fmt.Println(marks)
	fmt.Println(passed)
}
