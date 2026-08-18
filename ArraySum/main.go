package main

import "fmt"

func main() {
	//Find the sum of only the even numbers.

	nums := []int{12, 25, 7, 30, 18, 9, 14}
	sum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			sum = sum + nums[i]
		}
	}
	fmt.Println(sum)

}
