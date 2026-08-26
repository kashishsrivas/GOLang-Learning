package main

import "fmt"

func main() {
	//reverse a string using string manipulation + two pointers.
	str := "hello"

	//swap characters, we'll convert the string into a []byte
	chars := []byte(str)

	//pointers
	left := 0
	right := len(chars) - 1

	//left = 0, right = 4 → 0 < 4 ✅
	for left < right {
		//Put the right character on the left, and the left character on the right.
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
	fmt.Println(string(chars))
}
