package main

import "fmt"

func main() {
	//Find the character with the highest frequency.

	str := "aabbbcc"

	frequency := make(map[byte]int)

	max := 0
	maxChar := byte(0)

	for i := 0; i < len(str); i++ {
		frequency[str[i]]++
	}

	for ch, count := range frequency {
		if count > max {
			max = count
			maxChar = ch
		}
	}

	fmt.Printf("%c", maxChar)
}
