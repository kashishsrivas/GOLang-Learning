package main

import "fmt"

func main() {
	//count of numbers from 1 to n that are divisible by 3.
	// n := 20
	// count := 0

	// for i := 1; i <= n; i++ {
	// 	if i%3 == 0 {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	//    The count of numbers from 1 to n divisible by 3. and The sum of those numbers.
	n := 20
	count := 0
	total := 0

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			total += i
			count++
		}
	}
	fmt.Println(total)
	fmt.Println(count)
}
