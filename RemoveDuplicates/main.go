package main

import "fmt"

func main() {
	//remove duplicates from a sorted slice.

	nums := []int{1, 1, 2, 2, 3}

	slow := 0

	for fast := 1; fast < len(nums); fast++ {

		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	fmt.Println(nums[:slow+1])
}
