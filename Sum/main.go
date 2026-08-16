package main

import "fmt"

func main() {
	// total := 0
	// for i := 1; i <= 10; i++ {
	// 	total = total + i
	// }
	// fmt.Println(total)

	//sum of all even numbers
	// total := 0
	// for i := 2; i <= 20; i += 2 {
	// 	total = total + i
	// }
	// fmt.Println(total)

	//sum of all even numbers from 1 to n.
	// n := 20
	// total := 0
	// for i := 2; i <= n; i += 2 {
	// 	total = total + i
	// }
	// fmt.Println(total)

	//sum of all numbers from 1 to n that are divisible by 3.
	n := 20
	total := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			total = total + i
		}
	}
	fmt.Println(total)
}
