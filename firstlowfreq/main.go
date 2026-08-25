package main

import "fmt"

//Find the first character with the lowest frequency.

func main() {
	str := "aabbbccde"

	frequency := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		frequency[str[i]]++
	}

	min := len(str)

	for _, count := range frequency {
		if count < min {
			min = count
		}
	}

	for i := 0; i < len(str); i++ {
		if frequency[str[i]] == min {
			fmt.Printf("%c", str[i])
			break
		}
	}
}
