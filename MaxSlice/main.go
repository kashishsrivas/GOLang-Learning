package main

import "fmt"

func main() {
	//Find the maximum/largestnum in a slice
	nums := []int{12, 25, 7, 30, 18}
	largestnum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > largestnum {
			largestnum = nums[i]
		}
	}
	fmt.Println(largestnum)

}
