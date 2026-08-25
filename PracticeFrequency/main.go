package main

import "fmt"

func main() {
	//find the frequency of every character.
	str := "banana"

	frequency := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		frequency[str[i]]++
	}

	for x, y := range frequency {
		fmt.Printf("%c: %d\n", x, y)
	}
}
