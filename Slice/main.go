package main

import (
	"fmt"
)

func main() {
	nums := []int{12, 25, 7, 30, 18}

	//Write a program using a for loop that prints every element.

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	//sum of all elements.
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	fmt.Println(total)
}
