package main

import "fmt"

func main() {

	//Find the first number divisible by 7
	// n := 50
	// for i := 1; i <= n; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }

	//Find the first number between 1 and 50 that is divisible by both 3 and 5
	n := 50
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
			break
		}
	}
}
