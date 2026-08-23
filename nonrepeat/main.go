package main

import "fmt"

//Find the first character that repeats in a string.

func main() {
	str := "abcdab"
	seen := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		if seen[str[i]] == true {
			fmt.Printf("%c", str[i])
			break
		}
		seen[str[i]] = true
	}
}
