package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}

	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
	fmt.Println(nums)
}
