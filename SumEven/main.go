package main

import "fmt"

func main() {
	//Find the sum of all positive even numbers.
	// nums := []int{12, -5, 7, -20, 18, -3, 14, -8}
	// sum := 0

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > 0 && nums[i]%2 == 0 {
	// 		sum = sum + nums[i]
	// 	}
	// }
	// fmt.Println(sum)

	//Find how many numbers are greater than 10 AND even.
	// nums := []int{12, 5, 18, 7, 24, 30, 9, 14}
	// count := 0

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > 10 && nums[i]%2 == 0 {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	//Find the largest even number that is greater than 10.
	// nums := []int{12, 5, 18, 7, 24, 30, 9, 14}
	// largeNum := nums[0]

	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] > 10 && nums[i]%2 == 0 {
	// 		if nums[i] > largeNum {
	// 			largeNum = nums[i]
	// 		}
	// 	}
	// }
	// fmt.Println(largeNum)

	//Find the smallest even number greater than 20.
	nums := []int{15, 22, 7, 34, 18, 9, 40, 13, 26}
	var smallestNum int
	found := false

	for i := 0; i < len(nums); i++ {
		if nums[i] > 20 && nums[i]%2 == 0 {
			if found == false {
				smallestNum = nums[i]
				found = true
			} else {
				if nums[i] < smallestNum {
					smallestNum = nums[i]
				}
			}
		}
	}
	fmt.Println(smallestNum)
}
