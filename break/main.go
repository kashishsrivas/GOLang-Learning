package main

import "fmt"

// break - Immediately exit the loop.
func main() {
	//Prints numbers from 1 to 10, but stops when it reaches 7.
	for i := 1; i <= 10; i++ {
		fmt.Println(i) //output will 1-7
		if i == 7 {
			break
		}
		// fmt.Println(i) // output will be 1-6
	}
}
