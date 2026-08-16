package main

import "fmt"

func main() {
	//smallest element in the slice.
	nums := []int{12, 25, 7, 30, 18}
	minNum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	fmt.Println(minNum)
}
