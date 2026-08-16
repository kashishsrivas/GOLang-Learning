package main

import "fmt"

// continue -  Skip the current iteration and move to the next one.
func main() {
	//Prints numbers from 1 to 10, but skips 5.

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
