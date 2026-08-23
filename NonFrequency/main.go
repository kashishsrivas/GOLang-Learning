package main

import "fmt"

func main() {
	//Find the first non-repeating character

	str := "aabbcde"

	frequency := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		frequency[str[i]]++
	}

	for i := 0; i < len(str); i++ {
		if frequency[str[i]] == 1 {
			fmt.Printf("%c", str[i])
			break
		}
	}
}
