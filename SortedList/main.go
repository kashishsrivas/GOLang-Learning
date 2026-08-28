package main

import "fmt"

func main() {
	//find two numbers in a sorted list that add up to a target.

	nums := []int{1, 2, 4, 7, 11, 15}
	target := 9

	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]

		if sum == target {
			fmt.Println(nums[left], nums[right])
			break
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}
