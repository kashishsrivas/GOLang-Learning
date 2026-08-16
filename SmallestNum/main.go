package main

import "fmt"

func main() {
	//Find the smallest number between 1 and 30 that is divisible by 4.
	// smallNum := 30

	// for i := 1; i <= smallNum; i++ {
	// 	if i%4 == 0 {
	// 		if i < smallNum {
	// 			smallNum = i
	// 		}
	// 	}
	// }
	// fmt.Println(smallNum)

	//Find the smallest odd number between 1 and 50 that is divisible by 3.
	smallNumber := 50

	for i := 1; i <= smallNumber; i++ {
		if i%3 == 0 && i%2 != 0 {
			if i < smallNumber {
				smallNumber = i
			}
		}
	}
	fmt.Println(smallNumber)
}
