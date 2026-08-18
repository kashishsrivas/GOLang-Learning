package main

import "fmt"

func main() {
	//How many positive even numbers and their sum
	// nums := []int{12, -5, 18, 7, -20, 24, 9, 30, -8, 14}

	// count := 0
	// sum := 0

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > 0 && nums[i]%2 == 0 {
	// 		count++
	// 		sum = sum + nums[i]
	// 	}
	// }
	// fmt.Println(count)
	// fmt.Println(sum)

	//Find the largest positive odd number
	// nums := []int{15, 8, -4, 22, 7, 31, 18, -10, 26, 5}
	// largestNum := nums[0]

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > 0 && nums[i]%2 != 0 {
	// 		if nums[i] > largestNum {
	// 			largestNum = nums[i]
	// 		}
	// 	}
	// }
	// fmt.Println(largestNum)

	//other way
	//Find the largest positive odd number.
	nums := []int{-12, 8, -4, -22, -6, -10}
	var largestNum int
	found := false

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i]%2 != 0 {
			if found == false {
				largestNum = nums[i]
				found = true
			} else {
				if nums[i] > largestNum {
					largestNum = nums[i]
				}
			}

		}
	}
	if found == true {
		fmt.Println(largestNum)
	} else {
		fmt.Println("no positive odd number")
	}
}
