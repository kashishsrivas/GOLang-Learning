package main

import "fmt"

func main() {
	//smallest element in the slice.
	// nums := []int{12, 25, 7, 30, 18}
	// minNum := nums[0]

	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] < minNum {
	// 		minNum = nums[i]
	// 	}
	// }
	// fmt.Println(minNum)

	//Count how many numbers in the slice are even.
	nums := []int{12, 25, 7, 30, 18}
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
