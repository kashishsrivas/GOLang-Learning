package main

import "fmt"

func main() {
	//Counts how many numbers from 1 to 50 are divisible by both 3 and 5.
	// count := 0
	// for i := 1; i <= 50; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println(i)
	// 		count++
	// 	}
	// }
	// fmt.Println(count) //tells us that how many numbers are divisible i.e 3

	//Counts how many even numbers AND how many odd numbers are between 1 and 20.
	evenCount := 0
	oddCount := 0

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Println("Even:", evenCount)
	fmt.Println("Odd:", oddCount)
}
