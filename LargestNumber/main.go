package main

import "fmt"

func main() {
	//Finds the largest number between 1 and 20 that is divisible by 3.
	// largeNum := 0

	// for i := 1; i <= 20; i++ {
	// 	if i%3 == 0 {
	// 		if i > largeNum {
	// 			largeNum = i
	// 		}
	// 	}
	// }
	// fmt.Println(largeNum)

	//Find the largest even number between 1 and 50 that is divisible by 3.
	// largeNumber := 0

	// for i := 1; i <= 50; i++ {
	// 	if i%3 == 0 && i%2 == 0 {
	// 		if i > largeNumber {
	// 			largeNumber = i
	// 		}
	// 	}
	// }
	// fmt.Println(largeNumber)

	//Find the largest number between 1 and 100 that is divisible by 4 but NOT divisible by 6.
	// largeNum := 0

	// for i := 1; i <= 100; i++ {
	// 	if i%4 == 0 && i%6 != 0 {
	// 		if i > largeNum {
	// 			largeNum = i
	// 		}
	// 	}
	// }
	// fmt.Println(largeNum)

	//Find the largest number ≤ 100 that is divisible by BOTH 3 and 4, but NOT divisible by 5.
	largeNum := 0

	for i := 1; i <= 100; i++ {
		if (i%3 == 0 && i%4 == 0) && i%5 != 0 {
			if i > largeNum {
				largeNum = i
			}
		}
	}
	fmt.Println(largeNum)
}
