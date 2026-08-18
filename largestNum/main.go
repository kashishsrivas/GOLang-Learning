package main

import "fmt"

// Find the largest even number in the slice.
func main() {
	// nums := []int{12, 25, 7, 30, 18, 9, 14}

	// largeNum := nums[0]

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i]%2 == 0 {
	// 		if nums[i] > largeNum {
	// 			largeNum = nums[i]
	// 		}
	// 	}
	// }
	// fmt.Println(largeNum)

	//Find the smallest positive even number.
	nums := []int{-12, 25, -7, 30, -18, 9, 14}
	var smallestNum int
	found := false //checks the condition if things are present or not.

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 && nums[i] > 0 {
			if found == false { //if present
				smallestNum = nums[i]
				found = true
			} else { //if not then compare and update
				if nums[i] < smallestNum {
					smallestNum = nums[i]
				}
			}
		}
	}
	fmt.Println(smallestNum)
}
