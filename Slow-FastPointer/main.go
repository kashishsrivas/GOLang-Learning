package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3, 4, 3}
	target := 3

	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != target {
			nums[slow] = nums[fast]
			slow++
		}
	}

	fmt.Println(nums[:slow])
}
