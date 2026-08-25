package main

import "fmt"

//Find the character with the lowest frequency.

func main() {
	str := "aabbbccde"

	frequency := make(map[byte]int)

	min := len(str)
	minChar := byte(0)

	for i := 0; i < len(str); i++ {
		frequency[str[i]]++
	}

	for ch, count := range frequency {
		if count < min {
			min = count
			minChar = ch
		}
	}
	fmt.Printf("%c", minChar)
}
