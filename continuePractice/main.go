package main

import "fmt"

func main() {
	//Print all numbers from 1 to 20, except the numbers divisible by 3.
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
