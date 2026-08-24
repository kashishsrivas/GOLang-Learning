package main

import (
	"fmt"
)

func main() {
	//Find the most frequent character
	//most repeated in short in this "b" is the output
	str := "aabbbcc"

	frequency := make(map[byte]int)
	//max     → highest count
	//maxChar → character having that count

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
